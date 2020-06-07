<template>
  <div>
    <v-card
      draggable
      @dragstart="stashTodo($event, todo)"
    >
      <v-card-text class="font-weight-bold">{{ todo.text }}</v-card-text>
      <v-card-subtitle class="d-flex">
        {{ todo.user.name }}
        <v-spacer></v-spacer>
        <v-icon @click="taskDialog = true">mdi-square-edit-outline</v-icon>
        <TaskDeleteBtn :todo-id="todo.id" />
      </v-card-subtitle>
    </v-card>
    <TaskForm
      :dialog="taskDialog"
      :todo="todo"
      mode="edit"
      @close="taskDialog = false"
    />
  </div>
</template>

<script>
import TaskForm from "~/components/TaskForm";
import TaskDeleteBtn from "~/components/TaskDeleteBtn";

export default {
  components: {
    TaskForm,
    TaskDeleteBtn
  },
  props: {
    todo: {
      type: Object,
      required: true
    },
  },
  data() {
    return {
      taskDialog: false
    };
  },
  methods: {
    stashTodo(event, todo) {
      event.dataTransfer.effectAllowed = "move";
      event.dataTransfer.dropEffect = "move";
      event.dataTransfer.setData("todoId", todo.id);
    }
  }
};
</script>
