import React, { CSSProperties, RefObject } from "react";
import M, { Modal } from "materialize-css";

export type ConfirmCallback = (payload?: Object) => void;

export type ModalPayload = {
	[key: string]: any
}

type ModalProps = {
	title?: string;
	onCancel?: Function;
	onConfirm?: ConfirmCallback;
}
type ModalState = {
	title?: string;
	body?: JSX.Element;
	instance: Modal | null
	payload?: ModalPayload
}

class ModalDialog extends React.Component {
	props: ModalProps;
	state: ModalState;
	modalRef: RefObject<HTMLDivElement>;

	constructor(props: ModalProps) {
		super(props);
		this.props = props;
		this.state = {instance: null, payload: {}, title: props.title};
		this.modalRef = React.createRef();
		this.open = this.open.bind(this);
		this.updatePayload = this.updatePayload.bind(this);
	}

	componentDidMount(): void {
		const instance: Modal = M.Modal.init(this.modalRef.current as unknown as MElements, {preventScrolling: false}) as unknown as Modal;
		this.setState({instance});
	}

	open(body?: JSX.Element, title?: string) {
		this.setState({body, title});
		if (this.state.instance) {
			this.state.instance.open();
		}
	}

	onCancelHandler() {
		if (this.props.onCancel) {
			this.props.onCancel(false);
		}
		this.state.instance?.close();
	}

	onConfirmHandler() {
		if (this.props.onConfirm) {
			this.props.onConfirm(this.state.payload);
		}
		this.state.instance?.close();
	}

	updatePayload(payload: ModalPayload) {
		this.setState({payload});
	}

	render() {
		return (
			<div ref={this.modalRef} className="modal" style={styleSheet}>
				<div className="p-4 black white-text">
					<h4 id="modal-question-title" className="mb-0">{this.state.title}</h4>
				</div>
				<div className="modal-content">
					<div id="modal-question-body">
						{this.state.body}
					</div>
				</div>
				<div className="modal-footer">
					<button id="btn-modal-reject" onClick={this.onCancelHandler.bind(this)}
							className="waves-green btn red" style={btnStyle}>Close
					</button>
					{this.props.onConfirm ? <button id="btn-modal-confirm" onClick={this.onConfirmHandler.bind(this)}
													className="waves-green btn" style={btnStyle}>Confirm
					</button> : ""}

				</div>
			</div>);
	}
}

const btnStyle: CSSProperties = {
	width: 100,
	marginLeft: 10,
};

const styleSheet: CSSProperties = {};
export default ModalDialog;
