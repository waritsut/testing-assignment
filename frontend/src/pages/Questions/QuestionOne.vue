<template>
  <div>
    <section>
      <base-card
        ><question-line></question-line>

        <method></method>
        <h4><u>ANS</u></h4>
        <answer></answer>
      </base-card>
    </section>
  </div>
</template>

<script>
import QuestionLine from '../../components/questionOnes/QuestionLine.vue';
import Method from '../../components/questionOnes/Method.vue';
import Answer from '../../components/questionOnes/Answer.vue';

export default {
  components: {
    QuestionLine,
    Method,
    Answer
  }
};
</script>

<style scoped>
u {
  border-bottom: 1px double;
}
</style>
