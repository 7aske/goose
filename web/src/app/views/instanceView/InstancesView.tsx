import * as React from "react";
import axios from "axios";
import InstanceItem from "./instanceItem/InstanceItem";
import InstanceType from "../../../@types/Intance";
import M, { Collapsible } from "materialize-css";
import { RefObject } from "react";

type InstancesViewProps = {};
type InstancesViewState = {
	instances: InstanceType[]
};

export default class InstancesView extends React.Component<InstancesViewProps, InstancesViewState> {
	ref: RefObject<HTMLUListElement>;

	constructor(props: InstancesViewProps) {
		super(props);
		this.state = {instances: []};
		this.ref = React.createRef();
		this.getInstances = this.getInstances.bind(this);
	}

	componentDidUpdate(prevProps: Readonly<InstancesViewProps>, prevState: Readonly<InstancesViewState>, snapshot?: any): void {
		const instances = M.Collapsible.init(this.ref?.current as unknown as MElements, {}) as unknown as Collapsible;
		instances.open(0);
	}

	componentDidMount(): void {
		this.getInstances();
	}

	getInstances() {
		axios.get("/api/app/search").then(res => {
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
		}).catch(err => console.log(err));
	}

	handleRefresh() {
		this.getInstances();
	}

	render() {
		return (
			<ul ref={this.ref} className="collapsible">
				{this.state.instances.map((inst, i) => <InstanceItem
					triggerRefresh={this.handleRefresh.bind(this)}
					key={i} inst={inst}
					running={inst.pid !== undefined}/>)}
			</ul>
		);
	};
};
