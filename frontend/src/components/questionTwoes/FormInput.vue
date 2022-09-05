<template>
  <div>
    Item Price :
    <input type="number" min="0" v-model.number="cashier.itemPrice" /> Received
    Cash :
    <input type="number" min="0" v-model.number="cashier.receivedCash" />

    <div class="container">
      1000฿ :
      <input
        type="number"
        min="0"
        v-model.number="cashier.cash.oneThousandNote"
      />&nbsp;&nbsp; &nbsp; 500฿ :
      <input
        type="number"
        min="0"
        v-model.number="cashier.cash.fiveHundredNote"
      />
      &nbsp;&nbsp; &nbsp; 100฿ :
      <input
        type="number"
        min="0"
        v-model.number="cashier.cash.oneHundredNote"
      />
      &nbsp;&nbsp; &nbsp; 50฿ :
      <input type="number" min="0" v-model.number="cashier.cash.fiftyNote" />
      &nbsp;&nbsp; &nbsp; 20฿ :
      <input type="number" min="0" v-model.number="cashier.cash.twentyNote" />
    </div>
    <div class="container">
      &nbsp;&nbsp;&nbsp; 10฿ :&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
      <input
        type="number"
        min="0"
        v-model.number="cashier.cash.tenCoin"
      />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp; 5฿ :
      <input
        type="number"
        min="0"
        v-model.number="cashier.cash.fiveCoin"
      />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp; 1฿ :
      <input type="number" min="0" v-model.number="cashier.cash.oneCoin" />
      &nbsp; 0.25฿ :
      <input
        type="number"
        min="0"
        v-model.number="cashier.cash.twentyFiveSatang"
      />

      <button class="cal-btn" type="button" @click="clickCal()">
        CALCULATE
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      cashier: {
        itemPrice: 0,
        receivedCash: 0,
        cash: {
          twentyFiveSatang: 0,
          oneCoin: 0,
          fiveCoin: 0,
          tenCoin: 0,
          twentyNote: 0,
          fiftyNote: 0,
          oneHundredNote: 0,
          fiveHundredNote: 0,
          oneThousandNote: 0
        }
      }
    };
  },

  methods: {
    clickCal() {
      this.$store.dispatch('questionTwoes/getTheChange', this.cashier);
    }
  }
};
</script>

<style scoped>
.container {
  margin-top: 20px;
  text-align: center;
}

input,
textarea {
  width: 70px;
  border: 2px solid rgb(116, 111, 111);
  font: inherit;
}

.cal-btn {
  margin-left: 2.5rem;
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
