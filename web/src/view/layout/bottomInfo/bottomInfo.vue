<template>
  <div class="bottom-info">
    <div>
      <!--
      <span>{{ $GIN_VUE_ADMIN.appName }}:</span>
      -->
      <span>
        <a href="/docs/proto/">Protocol Documentation</a>
      </span>
      <el-divider direction="vertical" />
      <span>
        <a href="/docs/swagger/index.html">RESTful API Documentation</a>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BottomInfo'
}
</script>

<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
