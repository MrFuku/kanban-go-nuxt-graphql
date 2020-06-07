<template>
  <v-dialog
    v-model="dialog"
    max-width="290"
  >
    <template v-slot:activator="{ on }">
      <v-icon v-on="on">mdi-delete</v-icon>
    </template>
    <v-card>
      <v-card-title class="headline">タスクの削除</v-card-title>
      <v-card-text>本当に削除しますか？</v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="dialog = false">キャンセル</v-btn>
        <v-btn
          color="error"
          @click="deleteTodo"
        >削除</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script>
import todos from "~/apollo/queries/todos.gql";
import deleteTodo from "~/apollo/mutations/deleteTodo.gql";

export default {
  props: {
    todoId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      dialog: false
    };
  },
  methods: {
    deleteTodo() {
      this.$apollo
        .mutate({
          mutation: deleteTodo,
          variables: {
            id: this.todoId
          },
          refetchQueries: [
            {
              query: todos
            }
          ]
        })
        .then(res => {
          this.dialog = false;
        })
        .catch(err => {
          console.log(err)
        });
    }
  }
};
</script>
