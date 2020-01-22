import * as React from "react";
import { getBackendIcon } from "../../../../utils/InstanceUtils";

export default class BackendIcon extends React.Component<any, any> {
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
