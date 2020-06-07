<template>
  <v-app id="inspire">
    <Header
      :drawer="drawer"
      @change="drawer = !drawer"
    >
      <SideNav />
    </Header>
    <v-content>
      <v-container
        fluid
        fill-height
      >
        <v-row style="height: 90%;">
          <v-col cols="6">
            <Board
              :todos="unfinishedTodos"
              :status="false"
              @updateTodo="updateTodo"
            >
              <h3>未完了</h3>
            </Board>
          </v-col>
          <v-col cols="6">
            <Board
              :todos="finishedTodos"
              :status="true"
              @updateTodo="updateTodo"
            >
              <h3>完了</h3>
            </Board>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import Header from "~/components/Header";
import SideNav from "~/components/SideNav";
import Board from "~/components/Board";
import todos from "~/apollo/queries/todos.gql";
import users from "~/apollo/queries/users.gql";
import updateTodo from "~/apollo/mutations/updateTodo.gql";

export default {
  components: {
    Header,
    SideNav,
    Board
  },
  data() {
    return {
      drawer: true
    };
  },
  computed: {
    unfinishedTodos() {
      return this.todos.filter(t => !t.done);
    },
    finishedTodos() {
      return this.todos.filter(t => t.done);
    }
  },
  methods: {
    updateTodo(todoId, status) {
      let todo = this.todos.find(t => t.id === todoId);
      console.log(todoId, status)
      if (!todo || todo.done === status) return;

      const { id, text, userId } = todo;
      const done = status;
      this.$apollo
        .mutate({
          mutation: updateTodo,
          variables: {
            id,
            text,
            done,
            userId
          },
          refetchQueries: [
            {
              query: todos
            }
          ]
        })
        .catch(err => {
          console.log(err);
        });
    }
  },
  apollo: {
    todos: {
      prefetch: true,
      query: todos
    },
    users: {
      prefetch: true,
      query: users
    }
  }
};
</script>
