<template>
    <div class="pa-3">
        <v-btn class="menu-button" tile depressed color="error" outlined>Friends</v-btn>
        <v-btn class="menu-button" tile depressed color="error" outlined>Events</v-btn>
        <v-btn class="menu-button" tile depressed color="error">Logout</v-btn>
    </div>
</template>

<script>
export default {
    name: 'ProfileMenu'
}
</script>

<style scoped>
.menu-button {
    width: 100%;
}
</style>