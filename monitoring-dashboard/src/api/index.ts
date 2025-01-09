import axios from 'axios';
import type {AxiosInstance} from 'axios';

// Axios 인스턴스 생성
const api: AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL, // .env 파일의 URL 사용
    timeout: 10000, // 요청 타임아웃 설정
});

// CPU 데이터를 가져오는 함수
export const fetchCpuData = async (): Promise<Record<string, any>> => {
    try {
        const response = await api.get('/cpu');
        return response.data.data.used;
    } catch (error) {
        console.error('Error fetching CPU data:', error);
        throw error;
    }
};