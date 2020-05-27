import React from 'react';
import {grpc} from "@improbable-eng/grpc-web";
import {Presence, Push} from "../../index";
import Pbf from "pbf";
import {Message} from "./dto";

var encoder = new TextEncoder("utf-8");
var decoder = new TextDecoder("utf-8")
export class Main extends React.Component {
    constructor() {
        console.log(Push)
        console.log(Presence)
        super();
        this.subscription = null;
        this.heartbeat = null;
        this.state = {
            username: "",
            channel: "",
            writing: false,
            messages: [],
            online: []

        }
        const transport = grpc.FetchReadableStreamTransport()
        this.pushClient = new Push.Client("https://tawny.bobby-demo.site", { transport });
        this.presenceClient = new Presence.Client("https://tawny.bobby-demo.site", { transport });

    }

    async componentDidMount() {

        function getUrlVars() {
            var vars = {};
            window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function (m, key, value) {
                vars[key] = value;
            });
            return vars;
        }
        var vars = getUrlVars()
        if (vars['channel'] && vars['username']) {
            this.setState({...this.state, username:  vars['username'], channel:vars['channel']})
            this.start()
        }
    }
    start = async () => {
        console.log(this.pushClient)
        const subscribeInput = new Push.SubscribeInput();
        subscribeInput.setTopic('chat');
        subscribeInput.setChannel(this.state.channel);
        const stream = this.pushClient.subscribe(subscribeInput);
        stream.on('data', (owl) => {
            if (owl) {
                const data = owl.getData();
                var pbf = new Pbf(data);
                this.setState({ messages: [...this.state.messages, Message.read(pbf)] })
            }
        });
        stream.on('end', (status) => {
            if (status) {
                console.log(status)
            }
        });
        stream.on('status', (status) => {
            if (status) {
                console.log(status)
            }

        });


        const presenceInput = new Presence.PresenceSubscribeInput()
        presenceInput.setTopic('chat');
        presenceInput.setChannel(this.state.channel);
        console.log(this.presenceClient)
        const streamPresence = this.presenceClient.presenceSubscribe(presenceInput);
        streamPresence.on('data', (data) => {
            if (data) {
                const presenceList = data.getPresencesList()
                const activeUsers = presenceList.map((presence)=>{
                    return JSON.parse(decoder.decode(presence.getState()));
                })
                this.setState({...this.state, online:activeUsers })
            }
        });
        streamPresence.on('end', (status) => {
            if (status) {
                console.log(status)
            }
        });
        streamPresence.on('status', (status) => {
            if (status) {
                console.log(status)
            }
        });

        setInterval(()=>{
          this.heartBeat()
        }, 10000)
        this.heartBeat()
    }

    heartBeat = async()=>{
        const heartBeatInput = new Presence.HeartBeatInput()
        heartBeatInput.setTopic('chat');
        heartBeatInput.setChannel(this.state.channel);
        heartBeatInput.setState(encoder.encode(JSON.stringify({
            writing: this.state.writing,
            username:this.state.username
        })))
        heartBeatInput.setKey(this.state.username)
        this.presenceClient.heartBeat(heartBeatInput, (err) => {
            if (err) {
                console.log(err);
                return
            }
        })
    }
    handleLogin = async (event) => {
        event.preventDefault();
        await this.setState({ ...this.state, username: event.target.username.value, channel: event.target.channel.value, messages: [] })
        this.start()
    }

    handleSubmit = (event) => {
        event.preventDefault();
        const data = new FormData(event.target);
        var objData = Object.fromEntries(data);
        const pushInput = new Push.PushInput();
        pushInput.setTopic('chat');
        pushInput.setChannel(this.state.channel);
        var pbf = new Pbf();
        Message.write({ ...objData, username: this.state.username }, pbf);
        var buffer = pbf.finish();
        pushInput.setData(buffer);
        this.pushClient.publish(pushInput, (err) => {
            if (err) {
                console.log(err);
                return
            }
        });
    }
    triggerWriting = async (writing)=> {
       await this.setState({...this.state, writing})
       this.heartBeat()

    }

    render() {
        return (
            <div>
                {!this.state.username ? <form onSubmit={this.handleLogin}>
                        <h1>Please pick a username </h1>
                        <label>Username</label>
                        <input id="username" type="text"></input>
                        <label>Channel</label>
                        <input id="channel" type="text"></input>
                        <button>Login</button>
                    </form> :
                    <div><h2>
                        Loged {this.state.username} to {this.state.channel}
                    </h2>
                        <h4>
                            online ({this.state.online.length}): {this.state.online.map((u) => {
                            return <p>{u.username} {u.writing ? "..." : "connected"}</p>
                        })}
                        </h4>
                        <div>
                            {this.state.messages.map((m) => {
                                return <p>
                                    {m.username}: {m.message}
                                </p>
                            })}
                        </div>

                        <form onSubmit={this.handleSubmit}>
                            <label htmlFor="birthdate">Message</label>
                            <input onBlur={()=> {this.triggerWriting(false)}} onFocus={() => { this.triggerWriting(true) }} id="message" name="message" type="text" />
                            <button>Send !</button>
                        </form>

                    </div>}

            </div>
        );
    }
}