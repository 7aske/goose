import * as React from "react";
import axios from "axios";
import M from "materialize-css";


export default class LoginView extends React.Component<any, any> {
	userInputRef: React.RefObject<HTMLInputElement>;
	passInputRef: React.RefObject<HTMLInputElement>;

	constructor(props: any) {
		super(props);
		this.state = {user: "", pass: ""};
		this.userInputRef = React.createRef();
		this.passInputRef = React.createRef();
	}

	login(ev: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
		ev.preventDefault();
		axios.post("/auth", {user: this.state.user, pass: this.state.pass}).then(res => {
			M.toast({html: "login successful - redirecting", classes: "rounded green"});
			const token = res.data.token;
			sessionStorage.setItem("token", token);
			setTimeout(() => {
				window.location.href = "/";
			}, 2000);
		}).catch(err => {
			M.toast({html: "invalid credentials", classes: "rounded red"});
		});
	}

	onChange(ev: React.ChangeEvent) {
		if (ev.target.id === "user") {
			this.setState({user: (ev.target as HTMLInputElement).value});
		}
		if (ev.target.id === "pass") {
			this.setState({pass: (ev.target as HTMLInputElement).value});
		}
	}

	render() {
		return (
			<div className="row container">
				<form className="col s12">
					<div className="row">
						<div className="col m3 s12"/>
						<div className="input-field col m6 s12">
							<i className="material-icons prefix">account_circle</i>
							<input onChange={this.onChange.bind(this)} id="user" type="text" className="validate"/>
							<label htmlFor="user">Username</label>
						</div>
					</div>
					<div className="row">
						<div className="col m3 s12"/>
						<div className="input-field col m6 s12">
							<i className="material-icons prefix">lock</i>
							<input onChange={this.onChange.bind(this)} id="pass" type="tel" className="validate"/>
							<label htmlFor="pass">Password</label>
						</div>
					</div>
					<button onClick={this.login.bind(this)}
							className="waves-light btn cyan"><i
						className="material-icons right">arrow_right</i>Login
					</button>
				</form>
			</div>
		);
	};
};
