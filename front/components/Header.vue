<template>
  <div>
    <v-navigation-drawer
      :value="drawer"
      :clipped="$vuetify.breakpoint.lgAndUp"
      app
    >
      <slot></slot>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="$vuetify.breakpoint.lgAndUp"
      app
      color="blue darken-3"
      dark
    >
      <v-app-bar-nav-icon @click.prevent="$emit('change')"></v-app-bar-nav-icon>
      <v-toolbar-title class="ml-0 pl-4">
        <span class="hidden-sm-and-down">Kanban App</span>
      </v-toolbar-title>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  props: {
    drawer: {
      type: Boolean,
      required: true
    }
  }
};
</script>
