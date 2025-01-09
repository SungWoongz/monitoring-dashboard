# monitoring-dashboard

# VM Server Dashboard 기획안

## **1. 개요**

### **목적**
- 실시간으로 VM 서버의 메모리 사용량 데이터를 시각적으로 확인.
- 총 메모리, 사용 중인 메모리, 사용률, 캐시된 메모리 등 주요 정보를 차트와 텍스트로 제공.

### **기능**
1. 실시간 메모리 사용량 표시.
2. 주요 메모리 상태(총 메모리, 사용률, 사용 중/여유 메모리 등) 텍스트 요약.
3. 히스토리 그래프로 메모리 사용량 추이 확인.
4. 메모리 구성(Free, Cached, Active 등) 시각화.

---

## **2. 페이지 레이아웃**

### **1. 헤더**
- **페이지 제목**: "VM Server Dashboard"
- **업데이트 시간**: 최근 데이터 업데이트 시간을 표시.

### **2. 주요 정보 요약 카드**
- **구성**: 주요 정보를 카드 형식으로 표시.
- 예시:
  - 총 메모리: `8 GB`
  - 사용 메모리: `4 GB`
  - 사용률: `50.2%`
  - 캐시 메모리: `3.5 GB`

### **3. 실시간 메모리 사용량 차트**
- **구성**: 
  - 원형 차트(Pie Chart)로 Free/Used 비율 표시.
  - 색상 구분:
    - Used: 빨간색
    - Free: 녹색
  - 레이블로 퍼센트 표시.

### **4. 메모리 상세 데이터 테이블**
- **구성**:
  - 테이블로 세부 메모리 데이터를 표시.
  - 열 구성:
    - 항목: Total, Free, Cached, Active, Buffers, Dirty 등.
    - 값: 각각의 메모리 값.

### **5. 메모리 사용량 히스토리 그래프**
- **구성**:
  - 라인 차트(Line Chart)로 시간에 따른 메모리 사용량 추이 표시.
  - X축: 시간
  - Y축: 메모리 사용량 (GB 또는 MB 단위)

---

## **3. API 설계**

### **1. 백엔드 API**
- 데이터 제공 API:
  - **URL**: `/api/memory`
  - **Method**: `GET`
  - **Response** (JSON):
    ```json
    {
      "total": 8216653824,
      "free": 208629760,
      "used": 4124753920,
      "usedPercent": 50.199923,
      "active": 2115502080,
      "inactive": 5184315392,
      "cached": 3646935040,
      "buffers": 236335104,
      "dirty": 38293504
    }
    ```
- 데이터 갱신 주기:
  - 실시간 갱신이 필요하면 5초 간격으로 API 호출.

---

## **4. Vue 프론트엔드 구성**

### **1. 주요 기술 스택**
- **프레임워크**: Vue 3 (또는 Vue 2.7)
- **라이브러리**:
  - Axios: API 호출
  - Chart.js: 차트 시각화
  - Tailwind CSS: 스타일링

### **2. Vue 컴포넌트 구조**
1. **`App.vue`**: 전체 대시보드 레이아웃 및 데이터 관리.
2. **`Header.vue`**: 페이지 제목과 마지막 업데이트 시간 표시.
3. **`SummaryCards.vue`**: 주요 메모리 정보를 카드 형식으로 표시.
4. **`PieChart.vue`**: 메모리 사용량 비율을 시각화한 원형 차트.
5. **`MemoryTable.vue`**: 메모리 세부 데이터를 테이블로 표시.
6. **`LineChart.vue`**: 시간에 따른 메모리 사용량 추이를 표시하는 라인 차트.

---

### **3. Vue 파일별 기능**

#### **3-1. `App.vue`**
- **역할**: 전체 데이터 관리 및 컴포넌트 배치.
- **템플릿**:
  ```vue
  <template>
    <div class="app">
      <Header :lastUpdated="lastUpdated" />
      <SummaryCards :data="memoryData" />
      <PieChart :data="memoryData" />
      <MemoryTable :data="memoryData" />
      <LineChart :history="memoryHistory" />
    </div>
  </template>
  ```

#### **3-2. `Header.vue`**
- **역할**: 대시보드 제목과 데이터 갱신 시간 표시.
- **템플릿**:
  ```vue
  <template>
    <header>
      <h1>VM Server Dashboard</h1>
      <p>Last updated: {{ lastUpdated }}</p>
    </header>
  </template>
  ```

### **3-3. `SummaryCard.vue`**
- **역할**: 주요 정보를 카드 형태로 표시.
- **템플릿**:
```vue
<template>
  <div class="summary-cards">
    <div class="card" v-for="(value, key) in summary" :key="key">
      <h2>{{ key }}</h2>
      <p>{{ value }}</p>
    </div>
  </div>
</template>
```

### **3-4. `PieChart.vue`**
- **역할**: Free/Used 메모리를 시각화.
- **라이브러리**: Chart.js
- **템플릿**:
```vue
<template>
  <canvas id="memory-pie-chart"></canvas>
</template>
```

### **3-5. `LineChart.vue`**
- **역할**: 메모리 사용량 히스토리 표시.
- **템플릿**:
```vue
<template>
  <canvas id="memory-line-chart"></canvas>
</template>
```

#### **3-6. `MemoryTable.vue`**
- **역할**: 메모리 세부 데이터를 테이블로 표시.
- **템플릿**:
  ```vue
  <template>
    <table>
      <thead>
        <tr>
          <th>Metric</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(value, key) in memoryData" :key="key">
          <td>{{ key }}</td>
          <td>{{ value }}</td>
        </tr>
      </tbody>
    </table>
  </template>
  ```

---

## **5. UI 시안**

### **구성**
1. **헤더**
   - "VM Server Dashboard"라는 제목.
   - 마지막 데이터 업데이트 시간을 표시.

2. **카드 섹션**
   - 주요 메모리 상태를 카드로 표시.
     - 예: 총 메모리, 사용 메모리, 사용률, 캐시 메모리.

3. **원형 차트**
   - Free/Used 메모리 비율을 원형 차트로 시각화.

4. **테이블**
   - 메모리 세부 데이터를 테이블로 나열.
     - 각 항목과 값을 한 행에 표시.

5. **라인 차트**
   - 시간에 따른 메모리 사용량 추이를 표시.

---

## **6. 추가 기능 제안**

1. **메모리 경고 시스템**
   - 메모리 사용률이 특정 임계값(예: 80%)을 초과하면 경고 메시지 표시.

2. **데이터 갱신 주기 사용자 설정**
   - 데이터 갱신 주기를 사용자가 설정할 수 있도록 옵션 제공.

3. **반응형 UI**
   - 대시보드가 다양한 화면 크기(모바일, 태블릿, 데스크톱)에 최적화되도록 반응형 디자인 적용.

4. **테마 설정**
   - 다크 모드와 라이트 모드 지원.

---

## **7. 구현 순서**

1. **Vue 프로젝트 초기화**
   - Vue CLI 또는 Vite를 사용하여 프로젝트 생성.

2. **컴포넌트 개발**
   - 각 컴포넌트(`Header`, `SummaryCards`, `PieChart`, `MemoryTable`, `LineChart`)를 개별적으로 개발.

3. **API 연동**
   - 백엔드 `/api/memory` 엔드포인트와 Axios로 연결.

4. **Chart.js 통합**
   - `PieChart`와 `LineChart`에 Chart.js를 활용하여 시각화 구현.

5. **UI 스타일링**
   - Tailwind CSS 또는 기타 CSS 프레임워크로 UI 구성.

6. **실시간 데이터 갱신**
   - `setInterval` 또는 Vue의 `watch`를 사용하여 실시간 데이터 갱신 구현.

---

## **8. 기대 결과**

- 실시간으로 VM 서버의 메모리 상태를 모니터링할 수 있는 대시보드.
- 주요 정보를 직관적으로 확인할 수 있는 UI.
- 시간이 지남에 따라 메모리 사용 추이를 확인할 수 있는 그래프.
