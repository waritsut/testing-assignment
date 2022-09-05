<template>
  <span class="badge" :class="type">
    <strong>{{ text }} </strong>
  </span>
</template>

<script>
export default {
  props: ['text', 'type']
};
</script>

<style scoped>
.badge {
  margin: 0.75rem;
  line-height: 45px;
  padding: 2px;
  height: 50px;
  width: 50px;
  border-radius: 50%;
  display: inline-block;
  background-color: #c49792;
  color: #ffffff;
}

.layerOne {
  background-color: #ad91a3;
  color: #ffffff;
}

.layerTwo {
  background-color: #9d91a3;
  color: #ffffff;
}
</style>
