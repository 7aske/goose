import * as React from "react";

type NavbarProps = {};
type NavbarState = {};

export default class Navbar extends React.Component<NavbarProps, NavbarState> {
	// constructor(props: NavbarProps) {
	// 	super(props);
	// }

	render() {
		return (
			<nav>
				<div className="nav-wrapper">
					<a href="/" className="brand-logo"><i className="material-icons">cloud</i></a>
					<ul className="right hide-on-med-and-down">
						{/*<li><a href="#"><i className="material-icons">search</i></a></li>*/}
						{/*<li><a href="#"><i className="material-icons">view_module</i></a></li>*/}
						{/*<li><a href="#"><i className="material-icons">refresh</i></a></li>*/}
						{/*<li><a href="#"><i className="material-icons">more_vert</i></a></li>*/}
					</ul>
				</div>
			</nav>
		);
	};
};
