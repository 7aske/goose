import * as React from "react";
import InstanceType from "../../../../@types/Intance";
import axios from "axios";
import { getBackendIcon, uptimeStr } from "../../../../utils/InstanceUtils";
import ModalDialog, { ModalPayload } from "../../../components/modal/ModalDialog";
import { CSSProperties, FC, RefObject } from "react";
import SettingsBody from "../../../components/modal/SettingsBody";
import LogListBody from "../../../components/modal/LogListBody";

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
	settingsModalRef: RefObject<ModalDialog>;
	logModalRef: RefObject<ModalDialog>;
	logDetailModalRef: RefObject<ModalDialog>;
	refreshTimer: NodeJS.Timeout;

	constructor(props: InstanceItemProps) {
		super(props);
		this.state = {inst: props.inst, running: props.running};
		this.settingsModalRef = React.createRef();
		this.logModalRef = React.createRef();
		this.logDetailModalRef = React.createRef();
		this.refreshInstance = this.refreshInstance.bind(this);
		this.refreshTimer = setInterval(() => {
			this.refreshInstance();
		}, 10000);
	}

	componentWillUnmount(): void {
		clearTimeout(this.refreshTimer);
	}

	instanceRun() {
		axios.put("/api/app/run", {name: this.state.inst.name}).then(res => {
			if (res.status === 200) {
				this.setState({running: true});
				M.toast({html: "instance started", classes: "rounded green"});
			}
		}).catch(err => {
			M.toast({html: err.response.data.message, classes: "rounded red"});
		});
	}

	instanceUpdate() {
		M.toast({html: "update started", classes: "rounded cyan"});
		axios.put("/api/app/update", {name: this.state.inst.name}).then(res => {
			if (res.status === 200) {
				this.setState({running: false});
				M.toast({html: res.data.message, classes: "rounded green"});
			}
		}).catch(err => {
			M.toast({html: err.response.data.message, classes: "rounded red"});
		});
	}

	instanceKill() {
		axios.put("/api/app/kill", {name: this.state.inst.name}).then(res => {
			if (res.status === 200) {
				this.setState({running: false});
				M.toast({html: res.data.message, classes: "rounded green"});
			}
		}).catch(err => {
			M.toast({html: err.response.data.message, classes: "rounded red"});
		});
	}

	instanceRemove() {
		axios.delete("/api/app/remove", {data: {name: this.state.inst.name}}).then(res => {
			if (res.status === 200) {
				this.setState({running: false});
				M.toast({html: res.data.message, classes: "rounded green"});
				this.props.triggerRefresh();
			}
		}).catch(err => {
			M.toast({html: err.response.data.message, classes: "rounded red"});
		});
	}

	async instanceSettings(payload: ModalPayload | undefined) {
		try {
			const res = await axios.put("/api/app/settings", {name: this.state.inst.name, settings: payload});
			if (res.status === 200) {
				M.toast({html: res.data.message, classes: "rounded green"});
				this.setState({inst: res.data.instance});
				this.refreshInstance();
			}
		} catch (err) {
			M.toast({html: err.response.data.message, classes: "rounded red"});
		}
	}

	openInstanceSettingsModal() {
		if (this.settingsModalRef.current) {
			const comp = <SettingsBody
				updatePayloadHandler={this.settingsModalRef.current.updatePayload}
				fields={[{
					name: "port",
					value: this.state.inst.port,
					type: "number",
					display_name: "Port",
				}, {
					name: "hostname",
					value: this.state.inst.hostname,
					type: "text",
					display_name: "Host",
				}, {
					name: "backend",
					value: this.state.inst.backend,
					type: "text",
					display_name: "Backend",
				}]}/>;
			this.settingsModalRef.current.open(comp);
		}
	}

	openDetailedModal(name: string, type: string) {
		if (this.logDetailModalRef.current) {
			axios.get(`/api/app/log?instance=${name}&type=${type}`).then(res => {
				if (res.status === 200) {
					const data = res.data.content;
					console.log(data);
					this.logDetailModalRef.current?.open(<div>
						<pre style={{whiteSpace: "pre-wrap"}} className="left-align">{data}</pre>
					</div>, `${type} - ${name}`);
				}
			}).catch(err => {
				console.error(err);
			});
		}
	}

	openInstanceLogsModal() {
		if (this.logModalRef.current) {
			const comp = <LogListBody openDetailedModal={this.openDetailedModal.bind(this)}
									  instanceName={this.state.inst.name}/>;
			this.logModalRef.current.open(comp);
		}
	}


	refreshInstance() {
		axios.get("/api/app/search?query=" + this.state.inst.id).then(res => {
			if (res.status === 200) {
				this.setState({inst: res.data.instance, running: res.data.running});
			}
		}).catch(err => {
			M.toast({html: "unable to fetch instance data", classes: "rounded red"});
		});
	}

	render() {
		return (
			<li>
				<ModalDialog ref={this.settingsModalRef} title={this.state.inst.name}
							 onConfirm={this.instanceSettings.bind(this)}/>
				<ModalDialog ref={this.logModalRef} title={this.state.inst.name + " Logs"}/>
				<ModalDialog ref={this.logDetailModalRef}/>
				<div className="collapsible-header font-weight-bold"><i
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
								{this.state.running ?
									<InstanceItemRow name={"Uptime"} val={this.state.inst.uptime}/> : ""}
								{this.state.running ?
									<InstanceItemRow name={"PID"} val={this.state.inst.pid}/> : ""}
								<InstanceItemRow name={"Backend"} val={this.state.inst.backend}/>
							</ul>
						</div>
					</div>
					<div className="row">
						{this.state.running ?
							<div>
								<button onClick={this.instanceKill.bind(this)} style={btnStyles}
										className="waves-light btn red darken-4 ml-2 mr-2"><i
									className="material-icons right">close</i>Kill
								</button>
								<button onClick={this.openInstanceLogsModal.bind(this)} style={btnStyles}
										className="waves-light btn cyan lighten-2 btn ml-2 mr-2"><i
									className="material-icons right">info_outline</i>Logs
								</button>
							</div>
							:
							<div>
								<button onClick={this.instanceRun.bind(this)} style={btnStyles}
										className="waves-light btn ml-2 mr-2"><i
									className="material-icons right">directions_run</i>Run
								</button>
								<button onClick={this.instanceUpdate.bind(this)} style={btnStyles}
										className="waves-light btn blue btn ml-2 mr-2"><i
									className="material-icons right">sync</i>Update
								</button>

								<button onClick={this.openInstanceSettingsModal.bind(this)} style={btnStyles}
										className="waves-light btn orange btn ml-2 mr-2"><i
									className="material-icons right">settings</i>Settings
								</button>

								<button onClick={this.instanceRemove.bind(this)} style={btnStyles}
										className="waves-light btn red btn ml-2 mr-2"><i
									className="material-icons right">delete_forever</i>Remove
								</button>
								<button onClick={this.openInstanceLogsModal.bind(this)} style={btnStyles}
										className="waves-light btn cyan lighten-2 btn ml-2 mr-2"><i
									className="material-icons right">info_outline</i>Logs
								</button>
							</div>
						}
					</div>
				</div>
			</li>
		);
	};
};

const btnStyles: CSSProperties = {
	width: 140,
};

class BackendIcon extends React.Component<any, any> {
	backendIconRef: React.RefObject<HTMLImageElement>;

	constructor(props: any) {
		super(props);
		this.backendIconRef = React.createRef();
	}

	componentDidMount(): void {
		if (this.backendIconRef.current) {
			console.log(this.backendIconRef.current);
			M.Tooltip.init(this.backendIconRef.current, {});
		}
	}

	render() {
		return (
			<img ref={this.backendIconRef} className="tooltipped" data-position="top" data-tooltip={this.props.name}
				 style={{height: 50, width: 50}}
				 alt={this.props.name} src={getBackendIcon(this.props.name)}/>);
	};
};

function InstanceLink(props: any) {
	let val = "http://" + (props.href as string).replace("https://", "");
	return <a rel="noopener noreferrer" target="blank" href={val}>{val}</a>;
}


class InstanceItemRow extends React.Component<any, any> {
	render() {
		let val;
		if (this.props.name === "Host" || this.props.name === "Repo") {
			val = <InstanceLink href={this.props.val}/>;
		} else if (this.props.name === "Uptime") {
			val = uptimeStr(this.props.val as number);
		} else if (this.props.name === "Backend") {
			val = <BackendIcon name={this.props.val}/>;
		} else if (this.props.name === "Run" || this.props.name === "Updated" || this.props.name === "Deployed") {
			val = new Date(this.props.val).toLocaleString("en-GB", {
				day: "numeric",
				month: "numeric",
				year: "numeric",
				hour: "numeric",
				minute: "numeric",
				second: "numeric",
			});
		} else {
			val = this.props.val;
		}


		return (
			<li className="collection-item">
				<div className="row mb-0">
					<div className="col s3 left-align font-weight-bold">
						{this.props.name}:
					</div>
					<div className="col s9 right-align truncate">
						{val}
					</div>
				</div>
			</li>
		);
	};
};
