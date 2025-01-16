// [utils/memUtils.ts]

// 바이트를 GiB로 변환하는 함수
const bytesToGiB = (bytes: number): number => {
    return parseFloat((bytes / (1024 ** 3)).toFixed(2)); // 소수점 2자리 고정
};

// Memory 데이터를 가공하는 함수
export const processMemData = (rawData: any[]): { total: number; used: number; free: number; percentUsed: number } => {
    if (!Array.isArray(rawData) || rawData.length === 0) {
        throw new Error('Invalid or empty memory data');
    }

    const latestEntry = rawData[0]; // 가장 최신 데이터
    const total = bytesToGiB(latestEntry.total);
    const used = bytesToGiB(latestEntry.used);
    const free = Math.max(total - used, 0);
    const percentUsed = parseFloat(latestEntry.usedPercent.toFixed(2));

    return { total, used, free, percentUsed };
};