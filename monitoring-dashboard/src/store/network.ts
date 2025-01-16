import {defineStore} from 'pinia';

export const useNetworkStore = defineStore('network', {
    state: () => ({
        networkData: null as Record<string, any> | null,
    }),
    actions: {
        setNetworkData(data: Record<string, any>) {
            this.networkData = data;
        },
    },
});