import * as React from "react";
import InputField, { Field } from "./InputField";


type SettingsBodyProps = {
	fields: Field[];
	updatePayloadHandler: Function;
};

type SettingsBodyState = {
	fields: Field[];
	data: FieldKeyValue;
};

type FieldKeyValue = { [key: string]: any }


export default class SettingsBody extends React.Component<SettingsBodyProps, SettingsBodyState> {
	constructor(props: SettingsBodyProps) {
		super(props);
		let data: FieldKeyValue = {};
		props.fields.forEach(field => {
			data[field.name] = field.value;
		});
		this.state = {fields: props.fields, data: data};
		this.props.updatePayloadHandler(data);
	}

	handleFieldChange(kv: FieldKeyValue) {
		const data = this.state.data;
		data[kv.name] = kv.value;
		this.setState({data});
		this.props.updatePayloadHandler(data);
	}

	render() {
		return (
			<div>
				{this.state.fields.map((field, i) => {
					return <InputField
						onInputChange={this.handleFieldChange.bind(this)}
						key={i}
						field={field}/>;
				})}
			</div>
		);
	};
};
