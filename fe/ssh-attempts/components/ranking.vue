<script setup lang="ts">
import type { LoginAttempt } from '~/models/attempts';

interface RankingProps {
  attempts: LoginAttempt[]
  nodename: string
}

const props = defineProps<RankingProps>()

const totalAttempts = computed(
  () => props.attempts.reduce((p, c) => p + c.count, 0)
)
</script>

<template>
  <table>
    <caption>
      <span>Login attempts on node</span>
      <span>&nbsp;</span>
      <b>{{ nodename }}</b>
    </caption>
    <thead>
      <tr>
        <th scope="col">Rank</th>
        <th scope="col">Username</th>
        <th scope="col">Attempts</th>
      </tr>
    </thead>
    <tbody>
      <template v-for="(attempt, index) in attempts">
        <tr>
          <td>{{ index + 1 }}</td>
          <td>{{ attempt.username }}</td>
          <td>{{ attempt.count }}</td>
        </tr>
      </template>
    </tbody>
    <tfoot>
      <tr>
        <th scope="row">Total</th>
        <td colspan="2">{{ totalAttempts }}</td>
      </tr>
    </tfoot>
  </table>
</template>

<style></style>