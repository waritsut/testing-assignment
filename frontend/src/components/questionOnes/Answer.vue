<template>
  <div class="container">
    <h1>X = {{ ans.xValue }} ,Y = {{ ans.yValue }} ,Z = {{ ans.zValue }}</h1>
    <strong> Find the value by index : </strong>
    <input type="number" min="0" v-model.number="inputedIndex" />
    <button type="button" @click="clickCal()">Find</button>

    <h3>VALUE : {{ valueByIndex }}</h3>
  </div>
</template>

<script>
export default {
  data() {
    return {
      inputedIndex: 0
    };
  },

  computed: {
    ans() {
      return this.$store.getters['questionOnes/getAnswerValue'];
    },

    valueByIndex() {
      return this.$store.getters['questionOnes/getvalueByIndex'];
    }
  },

  methods: {
    clickCal() {
      this.$store.dispatch('questionOnes/getValueByIndex', this.inputedIndex);
    }
  },

  mounted() {
    this.$store.dispatch('questionOnes/getAns');
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
  height: 8rem;
}

input,
textarea {
  width: 100px;
  border: 2px solid rgb(116, 111, 111);
  font: inherit;
}

button,
a {
  text-decoration: none;
  padding: 0.2rem 1rem;
  font: inherit;
  background-color: #3a0061;
  border: 1px solid #3a0061;
  color: white;
  cursor: pointer;
  border-radius: 30px;
  margin-left: 0.5rem;
  display: inline-block;
}

a:hover,
a:active,
button:hover,
button:active {
  background-color: #270041;
  border-color: #270041;
}
</style>
