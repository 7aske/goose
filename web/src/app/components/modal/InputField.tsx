import * as React from "react";
import { ChangeEvent, RefObject } from "react";

export type Field = {
	icon?: string;
	display_name: string;
	type: string;
	name: string;
	value: string | number | Date | boolean;
}


type InputFieldProps = {
	field: Field, onInputChange: Function,
};
type InputFieldState = {
	name: string;
	value: any;
};


export default class InputField extends React.Component<InputFieldProps, InputFieldState> {
	ref: RefObject<HTMLInputElement>;

	constructor(props: any) {
		super(props);
		this.ref = React.createRef();
		this.state = {value: props.field.value, name: props.field.name};
	}

	componentDidMount(): void {
		if (this.ref.current) {
			this.ref.current.dispatchEvent(new Event("focus"));
		}
	}

	onChangeHandler(ev: ChangeEvent<HTMLInputElement>) {
		this.setState({value: (ev.target as HTMLInputElement).value});
	}

	onKeyUpHandler() {
		this.props.onInputChange({name: this.state.name, value: this.state.value});
	}

	render() {
		return (<div className="row">
			<div className="input-field col s12 m6">
				<i className="material-icons black-text prefix">{this.props.field.icon}</i>
				<input onChange={this.onChangeHandler.bind(this)}
					   ref={this.ref}
					   onKeyUp={this.onKeyUpHandler.bind(this)}
					   value={this.state.value}
					   id={this.props.field.name}
					   type={this.props.field.type}
					   className="validate"/>
				<label htmlFor={this.props.field.name}>{this.props.field.display_name}</label>
			</div>
		</div>);
	}
}
