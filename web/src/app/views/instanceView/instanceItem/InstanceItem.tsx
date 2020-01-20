import * as React from "react";
import InstanceType from "../../../../@types/Intance";
import axios from "axios";
import "./InstanceItem.css";

type InstanceItemProps = {
	triggerRefresh: Function;
	inst: InstanceType;
	running: boolean;
};
type instanceItemState = {
	inst: InstanceType;
	running: boolean;
};

export default class instanceItem extends React.Component<InstanceItemProps, instanceItemState> {
	constructor(props: InstanceItemProps) {
		super(props);
		this.state = {inst: props.inst, running: props.running};
		this.refreshInstance = this.refreshInstance.bind(this);
		setInterval(() => {
			this.refreshInstance();
		}, 10000);
	}

	instanceRun() {
		axios.put("/api/app/run", {name: this.state.inst.name}).then(res => {
			if (res.status === 200) {
				this.setState({running: true});
				M.toast({html: "instance started", classes: "rounded"});
			}
		}).catch(err => console.error(err));
	}

	instanceUpdate() {
		axios.put("/api/app/update", {name: this.state.inst.name}).then(res => {
			if (res.status === 200) {
				this.setState({running: false});
				M.toast({html: res.data.message, classes: "rounded"});
			}
		}).catch(err => console.error(err));
	}

	instanceKill() {
		axios.put("/api/app/kill", {name: this.state.inst.name}).then(res => {
			if (res.status === 200) {
				this.setState({running: false});
				M.toast({html: res.data.message, classes: "rounded"});
			}
		}).catch(err => console.error(err));
	}

	instanceRemove() {
		axios.delete("/api/app/remove", {data: {name: this.state.inst.name}}).then(res => {
			if (res.status === 200) {
				this.setState({running: false});
				M.toast({html: res.data.message, classes: "rounded"});
				this.props.triggerRefresh();
			}
		}).catch(err => console.error(err));
	}

	refreshInstance() {
		axios.get("/api/app/search?query=" + this.state.inst.id).then(res => {
			if (res.status === 200) {
				this.setState({inst: res.data.instance, running: res.data.running});
			}
		}).catch(err => console.error(err));
	}

	render() {

		return (
			<li>
				<div className="collapsible-header"><i
					className={(this.state.running ? "green-text" : "red-text") + " material-icons"}>whatshot</i>{this.state.inst.name}
				</div>
				<div className="collapsible-body">
					<div className="row">
						<div className="col s6">
							<ul className="collection">
								<InstanceItemRow name={"ID"} val={this.state.inst.id}/>
								<InstanceItemRow name={"Name"} val={this.state.inst.name}/>
								<InstanceItemRow name={"Repo"} val={this.state.inst.repo}/>
								<InstanceItemRow name={"Root"} val={this.state.inst.root}/>
								<InstanceItemRow name={"Port"} val={this.state.inst.port}/>
								<InstanceItemRow name={"Host"} val={this.state.inst.hostname}/>
							</ul>
						</div>
						<div className="col s6">
							<ul className="collection">
								<InstanceItemRow name={"Deployed"} val={this.state.inst.deployed}/>
								<InstanceItemRow name={"Updated"} val={this.state.inst.last_updated}/>
								<InstanceItemRow name={"Run"} val={this.state.inst.last_run}/>
								<InstanceItemRow name={"Backend"} val={this.state.inst.backend}/>
								{this.state.running ?
									<InstanceItemRow name={"Uptime"} val={this.state.inst.uptime}/> : ""}
								{this.state.running ?
									<InstanceItemRow name={"PID"} val={this.state.inst.pid}/> : ""}
							</ul>
						</div>
					</div>
					<div className="row">
						{this.state.running ?
							<button onClick={this.instanceKill.bind(this)}
									className="waves-light btn red darken-4 ml-2 mr-2"><i
								className="material-icons right">close</i>Kill</button>
							:
							<div>
								<button onClick={this.instanceRun.bind(this)} className="waves-light btn ml-2 mr-2"><i
									className="material-icons right">directions_run</i>Run
								</button>
								<button onClick={this.instanceUpdate.bind(this)}
										className="waves-light btn blue btn ml-2 mr-2"><i
									className="material-icons right">sync</i>Update
								</button>

								<button className="waves-light btn orange btn ml-2 mr-2"><i
									className="material-icons right">settings</i>Settings
								</button>

								<button onClick={this.instanceRemove.bind(this)}
										className="waves-light btn red btn ml-2 mr-2"><i
									className="material-icons right">delete_forever</i>Remove
								</button>
							</div>
						}
					</div>
				</div>
			</li>
		);
	};
};

type InstanceItemRowProps = {
	name: string;
	val: string | number | Date;
};

type InstanceItemRowState = {
	name: string;
	val: string | number | Date;
};

export class InstanceItemRow extends React.Component<InstanceItemRowProps, InstanceItemRowState> {
	constructor(props: InstanceItemRowProps) {
		super(props);
		this.state = {name: props.name, val: props.val};
	}

	uptimeStr(uptime: number): string {
		let days = Math.floor(uptime / (1000 * 60 * 60 * 24));
		let hours = Math.floor((uptime % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
		let minutes = Math.floor((uptime % (1000 * 60 * 60)) / (1000 * 60));
		let seconds = Math.floor((uptime % (1000 * 60)) / 1000);
		return `${days}d ${hours}h ${minutes}m ${seconds}s`;
	}

	render() {
		let val;
		if (this.state.name === "Host" || this.state.name === "Repo") {
			val = "http://" + (this.state.val as string).replace("https://", "");
		} else if (this.state.name === "Uptime") {
			val = this.uptimeStr(this.state.val as number);
		} else {
			val = this.state.val;
		}


		return (
			<li className="collection-item">
				<div className="row mb-0">
					<div className="col s3 left-align" style={{fontWeight: "bold"}}>
						{this.state.name}:
					</div>
					<div className="col s9 right-align truncate">
						{
							this.state.name === "Host" || this.state.name === "Repo" ?
								<a target={"_blank"}
								   href={"" + val}>{"" + val}</a> : val
						}
					</div>
				</div>
			</li>
		);
	};
};
