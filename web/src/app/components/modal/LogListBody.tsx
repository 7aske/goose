import * as React from "react";
import axios from "axios";
import ModalDialog from "./ModalDialog";

type LogListBodyProps = {
	instanceName: string;
	openDetailedModal: Function;
};
type LogListBodyState = {
	logList: string[];
};

export default class LogListBody extends React.Component<LogListBodyProps, LogListBodyState> {
	modalRef: React.RefObject<ModalDialog>;

	constructor(props: LogListBodyProps) {
		super(props);
		this.state = {logList: []};
		this.modalRef = React.createRef();
	}

	componentDidMount(): void {
		axios.get(`/api/app/log?instance=${this.props.instanceName}`).then(res => {
			if (res.status === 200) {
				this.setState({logList: res.data.files});
			}
		}).catch(err => {
			console.error(err);
		});
	}

	openDetailedModal(name: string, type: string) {
		this.props.openDetailedModal(name, type);
	}

	render() {
		return (
			<div>
				<ul className="collection">
					{this.state.logList.map((item, i) => {
						return <LogListItem openDetailedModal={this.openDetailedModal.bind(this)}
											key={i}
											name={this.props.instanceName}
											type={item}/>;
					})}
				</ul>
			</div>
		);
	};
};

class LogListItem extends React.Component<any, any> {
	modalRef: React.RefObject<ModalDialog>;

	constructor(props: any) {
		super(props);
		this.modalRef = React.createRef();
	}

	openDetailedModal() {
		this.props.openDetailedModal(this.props.name, this.props.type);
	}

	render() {
		return (
			<li className="collection-item">
				<ModalDialog ref={this.modalRef} title={this.props.name + " " + this.props.type}/>
				<div onClick={this.openDetailedModal.bind(this)} style={{cursor: "pointer"}}
					 className="blue-text popout">
					{this.props.type}
				</div>
			</li>
		);
	}
}
