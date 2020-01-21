import React from "react";
import "materialize-css/dist/css/materialize.min.css";
import "materialize-css/dist/js/materialize.min";
import "./assets/stylesheets/materializeicons.css";
import "./assets/stylesheets/helpers.css";

import "./App.css";
import { BrowserRouter, Route } from "react-router-dom";
import InstancesView from "./app/views/instanceView/InstancesView";
import Navbar from "./app/components/navbar/Navbar";
import LoginView from "./app/views/loginView/LoginView";

class App extends React.Component {
	render() {
		return (
			<div className="App">
				<BrowserRouter>
					<Navbar/>
					<Route exact path="/" component={InstancesView}/>
					<Route exact path="/login" component={LoginView}/>
				</BrowserRouter>
			</div>
		);
	}
}


export default App;
