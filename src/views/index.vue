<template>
  <div class="dashboard">
    <!-- 左侧主内容区 -->
    <div class="main-content">
      <!-- 顶部欢迎区 -->
      <header class="header flex items-center">
        <div class="welcome">
          <h1>Hello, Margaret</h1>
          <p class="subtitle">Track team progress here. You almost reach a goal!</p>
        </div>
        <div class="date text text-gray-500">16 May, 2023</div>
      </header>

      <!-- 统计数据区 -->
      <div class="stats">
        <div class="stat-card flex items-center flex-row">
          <div class="stat-icon">👍</div>
          <div class="stat-info">
            <h3>Finished</h3>
            <div class="stat-value">18</div>
            <div class="stat-label">+8 tasks</div>
          </div>
        </div>

        <div class="stat-card">
          <div class="stat-icon">👍</div>
          <div class="stat-info">
            <h3>Finished</h3>
            <div class="stat-value">18</div>
            <div class="stat-label">+8 tasks</div>
          </div>
        </div>

        <div class="stat-card">
          <div class="stat-icon">👍</div>
          <div class="stat-info">
            <h3>Finished</h3>
            <div class="stat-value">18</div>
            <div class="stat-label">+8 tasks</div>
          </div>
        </div>
        <!-- 其他统计卡片类似 -->
      </div>

      <!-- 性能图表区 -->
      <div class="performance">
        <div class="section-header">
          <h2>Performance</h2>
          <select class="date-range">
            <option>01-07 May</option>
          </select>
        </div>
        <div class="chart">
          <!-- 这里放置图表组件 -->
        </div>
      </div>

      <!-- 当前任务列表 -->
      <div class="current-tasks">
        <div class="section-header">
          <h2>Current Tasks</h2>
          <div class="progress">Done 30%</div>
        </div>
        <div class="task-list">
          <div class="task-item">
            <div class="task-icon">📋</div>
            <div class="task-info">
              <h4>Product Review for UI8 Market</h4>
              <span class="status in-progress">In progress</span>
            </div>
            <div class="task-time">4h</div>
            <button class="more-options">...</button>
          </div>
          <!-- 其他任务项类似 -->
        </div>
      </div>
    </div>

    <!-- 右侧活动区 -->
    <div class="activity-sidebar">
      <div class="user-profile">
        <img src="http://gips2.baidu.com/it/u=195724436,3554684702&fm=3028&app=3028&f=JPEG&fmt=auto?w=1280&h=960"
          alt="Profile" class="avatar">
        <div class="user-info">
          <h3>Megan Norton</h3>
          <p>@megnorton</p>
        </div>
      </div>

      <div class="activity-feed">
        <h3>Activity</h3>
        <!-- 活动列表项 -->
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard {
  display: flex;
  gap: 24px;
  padding: 24px;
  background: #f5f6fa;
  font-family: 'Inter', sans-serif;
}

.main-content {
  flex: 1;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
}

.welcome h1 {
  font-size: 24px;
  margin-bottom: 8px;
}

.stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-bottom: 32px;
}

.stat-card {
  background: white;
  border-radius: 12px;
  padding: 16px;
  display: flex;
  align-items: center;
  gap: 12px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.stat-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f0f2f5;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  margin: 4px 0;
}

.performance {
  background: white;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.date-range {
  padding: 8px 16px;
  border-radius: 8px;
  border: 1px solid #e0e0e0;
  background: white;
}

.chart {
  height: 200px;
  /* 图表样式需要根据具体图表库调整 */
}

.current-tasks {
  background: white;
  border-radius: 12px;
  padding: 24px;
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.task-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border-radius: 8px;
  background: #f8f9fc;
}

.task-icon {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: white;
}

.task-info {
  flex: 1;
}

.status {
  font-size: 12px;
  padding: 4px 8px;
  border-radius: 12px;
}

.status.in-progress {
  background: #fff4e5;
  color: #ff9800;
}

.activity-sidebar {
  width: 300px;
  background: white;
  border-radius: 12px;
  padding: 16px;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 24px;
  border-bottom: 1px solid #eee;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
}

.user-info {
  flex: 1;
}

.actions {
  display: flex;
  gap: 8px;
}

.actions button {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: none;
  background: #f0f2f5;
  cursor: pointer;
}

.activity-feed {
  margin-top: 24px;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .dashboard {
    flex-direction: column;
  }

  .activity-sidebar {
    width: 100%;
  }

  .stats {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .stats {
    grid-template-columns: 1fr;
  }
}
</style>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      userName: 'Margaret',
      currentDate: new Date().toLocaleDateString('en-US', {
        day: 'numeric',
        month: 'short',
        year: 'numeric'
      })
    }
  }
}
</script>