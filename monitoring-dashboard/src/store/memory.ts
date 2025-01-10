import {defineStore} from 'pinia';

export const useMemoryStore = defineStore('memory', {
    state: () => ({
        memoryData: null as Record<string, any> | null,
    }),
    actions: {
        setMemoryData(data: Record<string, any>) {
            this.memoryData = data;
        },
    },
});