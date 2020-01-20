type InstanceType = {
	id: string;
	repo: string;
	name: string;
	root: string;
	port: number;
	hostname: string;
	deployed: Date;
	last_updated: Date;
	last_run: Date;
	uptime: number;
	backend: string;
	pid: number;
}


export default InstanceType;
