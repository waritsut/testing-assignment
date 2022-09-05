<template>
  <div class="container">
    <div class="elem">
      <base-badge
        v-for="(data, index) in dataSet"
        :key="index"
        :text="data"
      ></base-badge>
    </div>
    <div>
      <base-badge
        v-for="(elem, index) in layerOne"
        :key="index"
        :text="elem"
        type="layerOne"
      ></base-badge>
    </div>
    <div>
      <base-badge
        v-for="(elem, index) in layerTwo"
        :key="index"
        :text="elem"
        type="layerTwo"
      ></base-badge>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dataSet: [1, 'X', 8, 17, 'Y', 'Z', 78, 113],
      layerOne: ['+2', '+5', '+9', '+14', '+20', '+27', '+35'],
      layerTwo: ['+3', '+4', '+5', '+6', '+7', '+8']
    };
  }
};
</script>

<style scoped>
.container {
  width: 90%;
  text-align: center;
  padding: 1rem;
  margin: 2rem auto;
  margin-top: 0;
  border-style: solid;
  border-color: #d3d3d3;
  border-radius: 18px;
  max-width: 80rem;
  height: 15rem;
}
</style>
