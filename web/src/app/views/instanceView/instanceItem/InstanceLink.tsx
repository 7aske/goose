import * as React from "react";

export function InstanceLink(props: any) {
	let val = "http://" + (props.href as string).replace("https://", "");
	return <a rel="noopener noreferrer" target="blank" href={val}>{val}</a>;
}
