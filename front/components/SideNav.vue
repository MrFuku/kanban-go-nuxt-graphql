<template>
  <div>
    <v-list dense>
      <v-list-item @click="taskDialog = true">
        <v-icon>mdi-calendar-check</v-icon>
        タスク作成
      </v-list-item>
      <v-list-item @click="userDialog = true">
        <v-icon>mdi-account</v-icon>
        ユーザー作成
      </v-list-item>
    </v-list>
    <TaskForm
      :dialog="taskDialog"
      mode="new"
      @close="taskDialog = false"
    />
    <UserForm
      :dialog="userDialog"
      @close="userDialog = false"
    />
  </div>
</template>

<script>
import TaskForm from "~/components/TaskForm";
import UserForm from "~/components/UserForm";

export default {
  components: {
    TaskForm,
    UserForm
  },
  data() {
    return {
      taskDialog: false,
      userDialog: false
    };
  }
};
</script>
