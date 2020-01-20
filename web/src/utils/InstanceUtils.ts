import angular from "../assets/icons/angular-brands.svg";
import node from "../assets/icons/node-js-brands.svg";
import python from "../assets/icons/python-brands.svg";
import react from "../assets/icons/react-brands.svg";
import npm from "../assets/icons/npm-brands.svg";

export const getBackendIcon = (backend: string): string => {
	switch (backend) {
		case "npm":
			return npm;
		case "python":
		case "flask":
			return python;
		case "node":
			return node;
		case "react":
			return react;
		case "angular":
			return angular;
		default:
			return "";
	}
};


export const uptimeStr = (uptime: number): string => {
	let days = Math.floor(uptime / (1000 * 60 * 60 * 24));
	let hours = Math.floor((uptime % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
	let minutes = Math.floor((uptime % (1000 * 60 * 60)) / (1000 * 60));
	let seconds = Math.floor((uptime % (1000 * 60)) / 1000);
	return `${days}d ${hours}h ${minutes}m ${seconds}s`;
};
