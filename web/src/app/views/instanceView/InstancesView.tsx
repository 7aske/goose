import * as React from "react";
import axios from "axios";
import InstanceItem from "./instanceItem/InstanceItem";
import InstanceType from "../../../@types/Intance";
import M, { Collapsible } from "materialize-css";
import { RefObject } from "react";
import ModalDialog, { ModalPayload } from "../../components/modal/ModalDialog";
import SettingsBody from "../../components/modal/SettingsBody";
import {withRouter} from"react-router-dom";

type InstancesViewProps = {};
type InstancesViewState = {
	instances: InstanceType[]
};

class InstancesView extends React.Component<any, any> {
	ref: RefObject<HTMLUListElement>;
	deployModalRef: RefObject<ModalDialog>;
	instancesCollapsible?: Collapsible;
	state: InstancesViewState;
	constructor(props: InstancesViewProps) {
		super(props);
		this.state = {instances: []};
		this.ref = React.createRef();
		this.deployModalRef = React.createRef();
		this.getInstances = this.getInstances.bind(this);
		axios.interceptors.request.use(function (config) {
			const token = sessionStorage.getItem("token");
			config.headers["Authorization"] = "Bearer " + token;
			if (config.data) {
				config.data["token"] = token;
			} else {
				config.data = {token};
			}
			return config;
		}, function (error) {
			return Promise.reject(error);
		});
	}

	componentDidUpdate(prevProps: Readonly<InstancesViewProps>, prevState: Readonly<InstancesViewState>, snapshot?: any): void {
		this.instancesCollapsible = M.Collapsible.init(this.ref?.current as unknown as MElements, {}) as unknown as Collapsible;
	}

	componentDidMount(): void {
		this.getInstances();
		if (this.instancesCollapsible) {
			this.instancesCollapsible.open(0);
		}
	}

	getInstances() {
		axios.get("/api/app/search", {
			data: {
				token: sessionStorage.getItem("token"),
			}, headers: {
				"Authorization": "Bearer " + sessionStorage.getItem("token"),
			},
		}).then(res => {
			let data: InstanceType[] = res.data.running ? res.data.running : [];
			if (res.data.deployed) {
				(res.data.deployed as InstanceType[]).forEach(inst => {
					if (!data.find(i => i.id === inst.id)) {
						data.push(inst);
					}
				});
			}
			console.log(data);
			this.setState({
				instances: data,
			});
		}).catch(err => {
			console.dir(err);
			if (err.response.status === 401) {
				// window.location.href = "/login";
				this.props.history.push("/login");
			}
		});
	}

	handleRefresh() {
		this.getInstances();
	}

	openDeployDialog() {
		if (this.deployModalRef.current) {
			const comp = <SettingsBody fields={[{
				name: "repo",
				value: "",
				display_name: "Repo",
				type: "text",
			}, {
				name: "hostname",
				value: "",
				display_name: "Host",
				type: "text",
			}, {
				name: "backend",
				value: "",
				display_name: "Backend",
				type: "text",
			}]} updatePayloadHandler={this.deployModalRef.current.updatePayload}/>;
			this.deployModalRef.current.open(comp);
		}
	}

	instanceDeploy(payload?: ModalPayload): void {
		if (payload) {
			M.toast({html: "deploying instance", classes: "rounded cyan"});
			axios.post("/api/app/deploy", {
				repo: payload.repo,
				hostname: payload.hostname,
				backend: payload.backend,
			}).then(res => {
				if (res.status === 200) {
					M.toast({html: res.data.message, classes: "rounded green"});
					this.getInstances();
				}
			}).catch(err => {
				console.dir(err);
				if (err.response) {
					M.toast({html: err.response.data.message, classes: "rounded red"});
				}
				console.error(err);
			});
		}
	}

	render() {
		return (
			<div>
				<ModalDialog ref={this.deployModalRef} title="Deploy Instance"
							 onConfirm={this.instanceDeploy.bind(this)}/>
				<div className="p-3 left-align">
					<button onClick={this.openDeployDialog.bind(this)}
							className="waves-light btn cyan btn ml-2 mr-2"><i
						className="material-icons right">cloud_upload</i>Deploy
					</button>
				</div>
				<ul ref={this.ref} className="collapsible">
					{this.state.instances.map((inst, i) => <InstanceItem
						triggerRefresh={this.handleRefresh.bind(this)}
						key={i} inst={inst}
						running={inst.pid !== undefined}/>)}
				</ul>
			</div>
		);
	};
};

export default withRouter(InstancesView);
