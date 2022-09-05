<template>
  <div>
    <section>
      <base-card>
        <div>
          <cashier-table></cashier-table>
          <form-input></form-input>
        </div>
      </base-card>
    </section>
  </div>
</template>

<script>
import CashierTable from '../../components/questionTwoes/CashierTable.vue';
import FormInput from '../../components/questionTwoes/FormInput.vue';

export default {
  components: {
    CashierTable,
    FormInput
  }
};
</script>
