// [utils/dataProcessor.ts]
import { parseISO, format } from 'date-fns';

// CPU 데이터 처리
export const processCpuData = (rawData: any[]): { time: string; averageUsage: number }[] => {
    return rawData.map((entry) => {
        const usageArray = JSON.parse(entry.Usage);
        const averageUsage =
            usageArray.reduce((sum: number, value: number) => sum + value, 0) / usageArray.length;

        return {
            time: format(parseISO(entry.timestamp), 'HH:mm:ss'), // 시간 변환
            averageUsage: parseFloat(averageUsage.toFixed(2)), // 소수점 2자리 고정
        };
    });
};

// MEMORY 데이터 처리
export const processMemData = (rawData: any): { used: number; free: number } => {
    const latestData = rawData.data[0];
    const usedMemory = latestData.used;
    const totalMemory = latestData.total;
    const freeMemory = Math.max(totalMemory - usedMemory, 0); // 음수 방지

    return {
        used: usedMemory,
        free: freeMemory,
    };
};