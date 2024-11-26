<script setup lang="ts">
interface LoginAttemptsProps {
  nodename: string
}

const props = defineProps<LoginAttemptsProps>()

const { data, error } = await useLoginAttempts(props.nodename)

const totalAttempts = computed(
  () => {
    if (data.value) {
      return data.value.attempts.reduce((p, c) => p + c.count, 0)
    }

    return 0
  }
)
</script>

<template>
  <table class="ranking">
    <caption>
      <span>Login attempts on node</span>
      <span>&nbsp;</span>
      <b>{{ nodename }}</b>
    </caption>
    <thead>
      <tr>
        <th scope="col" class="header-cell">Rank</th>
        <th scope="col" class="header-cell">Username</th>
        <th scope="col" class="header-cell">Attempts</th>
      </tr>
    </thead>
    <tbody>
      <template v-if="data">
        <tr v-for="(attempt, index) in data.attempts">
          <td class="number-cell">{{ index + 1 }}</td>
          <td class="text-cell">{{ attempt.username }}</td>
          <td class="number-cell">{{ attempt.count }}</td>
        </tr>
      </template>
      <template v-if="error">
        <tr>
          <td colspan="3" class="text-cell">Could not retrieve information for this node</td>
        </tr>
      </template>
    </tbody>
    <tfoot>
      <tr>
        <th scope="row" class="header-cell">Total</th>
        <td colspan="2" class="number-cell">{{ totalAttempts }}</td>
      </tr>
      <tr>
        <td colspan="3">
          <span v-if="data">Data pulled at: {{ data.generatedAt }}</span>
        </td>
      </tr>
    </tfoot>
  </table>
</template>

<style>
.ranking {
  width: 100%;
  min-width: 500px;
  table-layout: auto;
  border-collapse: collapse;
}

.ranking th,
.ranking td {
  padding: 8px;
  border: 1px solid #ddd;
}

.ranking tbody tr:nth-child(odd) {
  background-color: #f9f9f9;
}

.ranking tbody tr:nth-child(even) {
  background-color: #e0e0e0;
}

.ranking caption {
  display: none;
  text-align: left;
}

.number-cell {
  text-align: right;
}

.text-cell {
  text-align: left;
}

.header-cell {
  text-align: center;
}
</style>