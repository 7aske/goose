import InstanceType from "../@types/Intance";

export const isRunning = (instance: InstanceType, running: InstanceType[]) => {
	return running.find((inst) => inst.id === instance.id) !== undefined ;
};
