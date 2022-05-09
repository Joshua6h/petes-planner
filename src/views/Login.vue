<template>
    <span>
        <h1 class="my-10">Login Page</h1>
        <v-container class="form pb-5">
            <login-form @raiseError="setError($event)"></login-form>
            <v-container class="error my-5" v-if="error != ''">
                <h5>{{error}}</h5>
            </v-container>
        </v-container>
    </span>
</template>

<script>
import LoginForm from '@/components/LoginForm.vue';

export default {
    name: 'Login',
    data() {
        return {
            error: ''
        }
    },
    components: {
        LoginForm
    },
    methods: {
        setError(err) {
            if(err) {
                this.error = err;
            }
            else{
                this.error = '';
            }
        }
    }
}
</script>

<style>
.form {
    background-color: #FCEDDA;
    max-width: 30%;
    padding: 3em 5em 3em 5em;
    outline: 1px solid #EE4E34;
    border-radius: 0.5em;
}
</style>