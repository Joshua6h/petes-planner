<template>
    <span>
        <h1 class='my-10'>Login Page</h1>
        <v-container class='form' style='background-color: #FCEDDA;'>
            <login-form></login-form>
        </v-container>
    </span>
</template>

<script>
import LoginForm from '@/components/LoginForm.vue';

export default {
    name: 'Login',
    components: {
        LoginForm
    }
}
</script>

<style scoped>
.form {
    max-width: 30%;
    padding: 3em 5em 3em 5em;
    outline: 1px solid #EE4E34;
    border-radius: 0.5em;
}
</style>