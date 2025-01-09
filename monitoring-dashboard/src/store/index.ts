// src/store/cpu.ts
import {defineStore} from 'pinia';

export const useCpuStore = defineStore('cpu', {
    state: () => ({
        cpuData: null as Record<string, any> | null,
    }),
    actions: {
        setCpuData(data: Record<string, any>) {
            console.log('Setting CPU data:', data); // 디버깅용 로그
            this.cpuData = data;
        },
    },
});