import * as React from "react";
import InstanceType from "../../../../@types/Intance";
import axios from "axios";
import ModalDialog, { ModalPayload } from "../../../components/modal/ModalDialog";
import { CSSProperties, RefObject } from "react";
import SettingsBody from "../../../components/modal/SettingsBody";
import LogListBody from "../../../components/modal/LogListBody";
import { InstanceItemRow } from "./InstanceItemRow";

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

	componentDidMount(): void {
		M.Tooltip.init(document.querySelectorAll(".tooltipped"), {});
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

	instanceSettings(payload: ModalPayload | undefined) {
		axios.put("/api/app/settings", {name: this.state.inst.name, settings: payload}).then(res => {
			if (res.status === 200) {
				M.toast({html: res.data.message, classes: "rounded green"});
				this.setState({inst: res.data.instance});
				this.refreshInstance();
			}
		}).catch(err => {
			M.toast({html: err.response.data.message, classes: "rounded red"});
		});
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
				<div className="collapsible-header font-weight-bold" style={{userSelect:"none"}}><i
					className={(this.state.running ? "black-text" : "black-text") + " material-icons"}>{this.state.running ? "cloud_done" : "cloud_off"}</i>{this.state.inst.name}
				</div>
				<div className="collapsible-body pl-1 pl-1">
					<div className="row">
						<div className="col s12 m6">
							<ul className="collection">
								<InstanceItemRow name={"ID"} val={this.state.inst.id}/>
								<InstanceItemRow name={"Name"} val={this.state.inst.name}/>
								<InstanceItemRow name={"Repo"} val={this.state.inst.repo}/>
								<InstanceItemRow name={"Root"} val={this.state.inst.root}/>
								<InstanceItemRow name={"Port"} val={this.state.inst.port}/>
								<InstanceItemRow name={"Host"} val={this.state.inst.hostname}/>
							</ul>
						</div>
						<div className="col s12 m6">
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
					<div className="pl-4 text-left center-on-small-only">
						<button disabled={this.state.running} onClick={this.instanceRun.bind(this)} style={btnStyles}
								data-position="top" data-tooltip="Run"
								className="btn-floating btn tooltipped black ml-1"><i
							className="material-icons right">directions_run</i>
						</button>
						<button disabled={this.state.running} onClick={this.instanceUpdate.bind(this)} style={btnStyles}
								data-position="top" data-tooltip="Update"
								className="btn-floating btn tooltipped black ml-1"><i
							className="material-icons right">sync</i>
						</button>
						<button disabled={this.state.running} onClick={this.openInstanceSettingsModal.bind(this)}
								style={btnStyles}
								data-position="top" data-tooltip="Settings"
								className="btn-floating btn tooltipped black ml-1"><i
							className="material-icons right">settings</i>
						</button>
						<button disabled={this.state.running} onClick={this.instanceRemove.bind(this)} style={btnStyles}
								data-position="top" data-tooltip="Remove"
								className="btn-floating btn tooltipped black ml-1"><i
							className="material-icons right">delete_forever</i>
						</button>
						<button disabled={!this.state.running} onClick={this.instanceKill.bind(this)} style={btnStyles}
								data-position="top" data-tooltip="Kill"
								className="btn-floating btn tooltipped black darken-4 ml-1"><i
							className="material-icons right">close</i>Kill
						</button>
						<button onClick={this.openInstanceLogsModal.bind(this)}
								style={btnStyles}
								data-position="top" data-tooltip="Logs"
								className="btn-floating btn tooltipped black lighten-2 ml-1"><i
							className="material-icons right">info_outline</i>
						</button>

					</div>
				</div>
			</li>
		);
	};
};

const btnStyles: CSSProperties = {
	// width: 140,
	marginBottom: 10,
};




