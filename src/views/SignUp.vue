<template>
    <span>
        <h1 class="my-10">Sign Up Page</h1>
        <v-container class="form">
            <sign-up-form @raiseError="setError($event)"></sign-up-form>
            <v-container class="error my-5" v-if="error != ''">
                <h5>{{error}}</h5>
            </v-container>
        </v-container>
    </span>
</template>

<script>
import SignUpForm from '@/components/SignUpForm.vue';

export default {
    name: 'SignUp',
    data() {
        return {
            error: ''
        }
    },
    components: {
        SignUpForm
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