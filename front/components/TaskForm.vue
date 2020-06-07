<template>
  <v-dialog
    :value="dialog"
    @input="$emit('close')"
    width="500"
  >
    <v-card>
      <v-card-title
        class="headline grey lighten-2"
        primary-title
      >
        {{ title }}
      </v-card-title>
      <v-container>
        <v-form
          ref="form"
          v-model="valid"
        >
          <v-textarea
            v-model="editTodo.text"
            outlined
            label="タスクの内容"
            :rules="textRules"
          />
          <v-row>
            <v-col cols="6">
              <v-select
                :items="todoStatuses"
                v-model="editTodo.done"
                outlined
                dense
                label="完了ステータス"
              />
            </v-col>
            <v-col cols="6">
              <v-select
                :items="users"
                item-text="name"
                item-value="id"
                v-model="editTodo.userId"
                outlined
                dense
                label="担当者"
                :rules="userIdRulues"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-container>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="$emit('close')">
          キャンセル
        </v-btn>
        <v-btn
          color="primary"
          @click="exec"
          :disabled="!valid"
        >
          {{ submitLabel }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import todos from "~/apollo/queries/todos.gql";
import users from "~/apollo/queries/users.gql";
import createTodo from "~/apollo/mutations/createTodo.gql";
import updateTodo from "~/apollo/mutations/updateTodo.gql";

export default {
  props: {
    dialog: {
      type: Boolean,
      required: true
    },
    mode: {
      type: String,
      required: true
    },
    todo: {
      type: Object,
      default: () => {
        return {
          text: "",
          userId: "",
          done: false
        };
      }
    }
  },
  data() {
    return {
      editTodo: {},
      todoStatuses: [
        { text: "未完了", value: false },
        { text: "完了", value: true }
      ],
      textRules: [
        v => !!v || "タスクの内容は必須です",
        v => (v && v.length <= 1000) || "1000文字以内で入力してください"
      ],
      userIdRulues: [v => !!v || "担当者は必須です"],
      valid: false
    };
  },
  methods: {
    createTodo() {
      const { text, done, userId } = this.editTodo;
      this.$apollo
        .mutate({
          mutation: createTodo,
          variables: {
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
        .then(res => {
          this.$emit("close");
        })
        .catch(err => {
          console.log(err);
        });
    },
    updateTodo() {
      const { id, text, done, userId } = this.editTodo;
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
        .then(res => {
          this.$emit("close");
        })
        .catch(err => {
          console.log(err);
        });
    }
  },
  computed: {
    title() {
      if (this.mode === "new") return "新規タスク作成";
      if (this.mode === "edit") return "タスク編集";
    },
    submitLabel() {
      if (this.mode === "new") return "作成";
      if (this.mode === "edit") return "更新";
    },
    exec() {
      if (this.mode === "new") return this.createTodo;
      if (this.mode === "edit") return this.updateTodo;
    }
  },
  watch: {
    // フォームの入力内容を初期化する
    // 雑なやり方しか思い浮かばず。他に良いやり方があれば教えてください
    dialog(newValue) {
      if (newValue) {
        this.editTodo = Object.assign({}, this.todo);
      } else {
        this.editTodo = Object.assign({}, this.todo);
        this.$refs.form.resetValidation();
      }
    }
  },
  apollo: {
    users: {
      query: users
    }
  }
};
</script>
