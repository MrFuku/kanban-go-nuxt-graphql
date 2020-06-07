<template>
  <v-card
    class="blue-grey lighten-4"
    height="100%"
    @drop="putTodo($event)"
    @dragover.prevent
    @dragenter.prevent
  >
    <v-container>
      <slot></slot>
      <v-row dense>
        <v-col
          v-for="(todo, i) in todos"
          :key="i"
          cols="12"
        >
          <TaskCard :todo="todo" />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import TaskCard from "~/components/TaskCard";

export default {
  components: {
    TaskCard
  },
  props: {
    todos: {
      type: Array,
      required: true
    },
    status: {
      type: Boolean,
      required: true
    }
  },
  methods: {
    putTodo(event) {
      const todoId = event.dataTransfer.getData("todoId");
      this.$emit("updateTodo", todoId, this.status);
    }
  }
};
</script>
