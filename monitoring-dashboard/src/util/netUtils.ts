// util/netUtils.ts
interface NetworkEntry {
    timestamp: string;
    bytesSent: number;
    bytesRecv: number;
}

interface ProcessedNetworkData {
    time: string;
    sent: number; // 전송량 (MB)
    recv: number; // 수신량 (MB)
}

export const processNetData = (data: NetworkEntry[]): ProcessedNetworkData[] => {
    return data.map((entry) => ({
        time: new Date(entry.timestamp).toLocaleTimeString(), // 시간 변환
        sent: parseFloat((entry.bytesSent / 1024 / 1024).toFixed(2)), // MB로 변환
        recv: parseFloat((entry.bytesRecv / 1024 / 1024).toFixed(2)), // MB로 변환
    }));
};