<template>
  <div>
    <div class="container">
      <h3 class="ques-line">Q2) CASHIER SYSTEM</h3>
      <div>
        <table class="data-table">
          <tr>
            <td><strong>Bank Note/Coin Value</strong></td>
            <td><strong>Amount</strong></td>
          </tr>
          <tr v-for="(value, key) in getCashInDrawer" :key="key">
            <td>{{ key }}฿</td>
            <td>{{ value }}</td>
          </tr>
        </table>
      </div>

      <div class="change-container">
        <h3>CHANGE: {{ getTheChange.change }}฿</h3>
        <table class="change-table">
          <tr>
            <th>1฿ :</th>
            <th>{{ getTheChange.changeCash.oneCoin }} unit</th>
          </tr>
          <tr>
            <th>5฿ :</th>
            <th>{{ getTheChange.changeCash.fiveCoin }} unit</th>
          </tr>

          <tr>
            <th>10฿ :</th>
            <th>{{ getTheChange.changeCash.tenCoin }}unit</th>
          </tr>

          <tr>
            <th>20฿ :</th>
            <th>{{ getTheChange.changeCash.twentyNote }} unit</th>
          </tr>

          <tr>
            <th>50฿ :</th>
            <th>{{ getTheChange.changeCash.fiftyNote }} unit</th>
          </tr>

          <tr>
            <th>100฿ :</th>
            <th>{{ getTheChange.changeCash.oneHundredNote }} unit</th>
          </tr>

          <tr>
            <th>500฿ :</th>
            <th>{{ getTheChange.changeCash.fiveHundredNote }}unit</th>
          </tr>

          <tr>
            <th>1000฿ :</th>
            <th>{{ getTheChange.changeCash.oneThousandNote }} unit</th>
          </tr>

          <tr>
            <th>0.25฿ :</th>
            <th>{{ getTheChange.changeCash.twentyFiveSatang }} unit</th>
          </tr>
        </table>
      </div>

      <div class="footer">
        <button type="button" @click="clickReset()">RESET</button>
        <strong> BALANCE : {{ getBalance.balance }}฿</strong>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      moneySet: ['1000', '500', '100', '50', '20', '10', '5', '1', '0.25']
    };
  },

  computed: {
    getBalance() {
      return this.$store.getters['questionTwoes/getBalance'];
    },

    getCashInDrawer() {
      const cash =
        this.$store.getters['questionTwoes/getBalance'].cashInCashier;

      return cash;
    },

    getTheChange() {
      return this.$store.getters['questionTwoes/getTheChange'];
    }
  },

  mounted() {
    this.$store.dispatch('questionTwoes/getBalance');
  },

  methods: {
    clickReset() {
      this.$store.dispatch('questionTwoes/resetCashier');
    }
  }
};
</script>

<style scoped>
.change-table {
  margin-left: auto;
  margin-right: auto;
}

.change-container {
  line-height: 1.6;
}
.footer {
  text-align: left;
  margin-left: 5rem;
}
.ques-line {
  text-align: left;
}
.data-table {
  border-collapse: collapse;
  width: 400px;
  float: left;
}
.data-table th,
.data-table td {
  padding: 5px;
  border: solid 1px #777;
}
.data-table th {
  background-color: lightblue;
}
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
  height: 25rem;
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
  margin-top: 0;
  margin-left: 0;
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
