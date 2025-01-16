import axios from 'axios';
import type {AxiosInstance} from 'axios';

// Axios 인스턴스 생성
const api: AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL, // .env 파일의 URL 사용
    timeout: 10000, // 요청 타임아웃 설정
});

// CPU 데이터 호출 (응답 원본 반환)
export const fetchCpuData = async (): Promise<any> => {
    try {
        const response = await api.get('/cpu');
        return response.data.data; // 원본 데이터 반환
    } catch (error) {
        console.error('Error fetching CPU data:', error);
        throw error;
    }
};

// MEMORY 데이터 가져오는 함수
// MEMORY 데이터 호출 (응답 원본 반환)
export const fetchMemData = async (): Promise<any> => {
    try {
        const response = await api.get('/mem');
        return response.data; // 원본 데이터 반환
    } catch (error) {
        console.error('Error fetching MEMORY data:', error);
        throw error;
    }
};

// NETWORK 데이터 호출
export const fetchNetData = async (): Promise<any> => {
    try {
        const response = await api.get('/net');
        return response.data.data;
    } catch (error) {
        console.error('Error fetching NETWORK data:', error);
        throw error;
    }
};

// Server Info 데이터 호출
export const fetchInfoData = async (): Promise<any> => {
    try {
        const response = await api.get('/info');
        return response.data.data;
    } catch (error) {
        console.error('Error fetching Server Info data:', error);
        throw error;
    }
};