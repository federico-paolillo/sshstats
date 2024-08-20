<script setup lang="ts">
import LoginAttempts from "~/components/login-attempts.vue";

const nodenames = useNodenames()
</script>

<template>
  <section>
    <h2>
      Login attempts
    </h2>
    <div class="login-attempts-container">
      <section class="login-attempt" v-for="nodename in nodenames">
        <h3>Node <em>{{ nodename }}</em></h3>
        <div class="table-scroll">
          <LoginAttempts :nodename></LoginAttempts>
        </div>
      </section>
    </div>
  </section>
</template>

<style>
.login-attempts-container {
  display: flex;
  flex-flow: row wrap;
  gap: 8px;
  justify-content: center;
}

.login-attempt {
  flex: 0 1;
  box-sizing: border-box;
  width: 100%;
  height: 100%;
}

.table-scroll {
  overflow-x: auto;
  width: 100%;
}
</style>