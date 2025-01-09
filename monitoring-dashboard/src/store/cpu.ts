import {defineStore} from 'pinia';

export const useCpuStore = defineStore('cpu', {
    state: () => ({
        cpuData: null as Record<string, any> | null,
    }),
    actions: {
        setCpuData(data: Record<string, any>) {
            this.cpuData = data;
        },
    },
});