import * as React from "react";
import { InstanceLink } from "./InstanceLink";
import { uptimeStr } from "../../../../utils/InstanceUtils";
import BackendIcon from "./BackendIcon";

export class InstanceItemRow extends React.Component<any, any> {
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
