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
        新規ユーザー作成
      </v-card-title>
      <v-form
        ref="form"
        v-model="valid"
      >
        <v-container>
          <v-text-field
            v-model="name"
            :rules="nameRulues"
            outlined
            label="ユーザー名"
          />
        </v-container>
        <v-divider />
        <v-card-actions>
          <v-spacer />
          <v-btn @click="$emit('close')">
            キャンセル
          </v-btn>
          <v-btn
            color="primary"
            @click="createUser"
            :disabled="!valid"
          >
            作成
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script>
import users from "~/apollo/queries/users.gql";
import createUser from "~/apollo/mutations/createUser.gql";

export default {
  props: {
    dialog: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      name: "",
      nameRulues: [
        v => !!v || "ユーザー名は必須です",
        v => (v && v.length <= 1000) || "1000文字以内で入力してください"
      ],
      valid: false
    };
  },
  methods: {
    createUser() {
      const name = this.name;
      this.$apollo
        .mutate({
          mutation: createUser,
          variables: {
            name
          },
          refetchQueries: [
            {
              query: users
            }
          ]
        })
        .then(res => {
          this.$emit('close');
        })
        .catch(err => {
          console.log(err)
        });
    }
  },
  watch: {
    // フォームの入力内容を初期化する
    // 雑なやり方しか思い浮かばず。他に良いやり方があれば教えてください
    dialog(newValue) {
      if (!newValue) {
        this.name = "";
        this.$refs.form.resetValidation();
      }
    }
  }
};
</script>
