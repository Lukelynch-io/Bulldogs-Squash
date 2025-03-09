<script setup lang="ts">
import { inject, onMounted, onUnmounted, ref, type Ref } from 'vue';
import type { PostData } from '@/datatypes/PostData';
import Post from '../components/PostComponent.vue';
import AlertBox from '@/components/AlertBox.vue';
import { MessageType } from '@/datatypes/MessageType';
import { GetPosts, GetPost } from '@/api_calls';
import Modal from '@/components/Modal.vue';
import { CurrentUser } from '@/auth/auth';

const currentUser = inject<Ref<CurrentUser | null>>("currentUser");
const postsArray = ref<PostData[]>([]); // Reactive array for blog posts
const errorOccurred = ref(false);
const errorMessage = ref('')
const errorTitle = ref('Error Title')
let postLoadInterval: ReturnType<typeof setInterval>;

async function RefreshBlogPosts() {
    // Fetch blog posts from the API
    const [returnArray, returnError] = await GetPosts();
    if (returnError != "") {
        errorOccurred.value = true;
        errorMessage.value = returnError
        return;
    }
    errorOccurred.value = false;
    postsArray.value = returnArray;
}

onMounted(() => {
    RefreshBlogPosts();
    postLoadInterval = setInterval(RefreshBlogPosts, 10000);
})

onUnmounted(() => {
    clearInterval(postLoadInterval)
})

const isPostModal = ref(false)
let postData: PostData
async function LoadPost(postId: string) {
    postData = await GetPost(postId)
    isPostModal.value = true;
}

function ClosePost() {
    isPostModal.value = false;
}

</script>

<template>
    <div class="post-collection-view">
        <div v-if="currentUser?.CanUserCreatePosts()" style="display:flex; justify-content:center">
            <RouterLink to="/Post">
                <button>Create new post</button>
            </RouterLink>
        </div>
        <div class="card-collection">
            <Post v-for="(post, index) in postsArray" :key=index v-bind="post" @expand-post="LoadPost" />
        </div>
        <AlertBox v-if="errorOccurred" :messageType="MessageType.Info" :messageTitle="errorTitle"
            :messageDescription="errorMessage" />
        <Transition>
            <Modal v-if="isPostModal" :closeModal="ClosePost" :custom-content-style="'margin: 5%; height: auto'">
                <div class="post-modal-grid">
                    <div id="post-modal-image">
                        <img :src="postData.imageUrl ?? ''" alt="Bulldog">
                    </div>
                    <div id="div-modal-text">
                        <h1>{{ postData.title }}</h1>
                        <p>{{ postData.description }}</p>
                    </div>
                </div>
            </Modal>
        </Transition>
    </div>
</template>

<style scoped>
.post-modal-grid {
    display: grid;
    grid-template:
        "a b";
}

.post-modal-grid p {
    text-align: left;
}

.post-collection-view {
    padding-left: 5%;
    padding-right: 5%;
    margin-top: 5rem;
}

.card-collection {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
    grid-template-rows: 100% repeat(auto-fit, minmax(150px, 1fr));
    gap: 1rem;
    margin-top: 1rem;
}

@media only screen and (min-width: 1200px) {
    .card-collection {
        grid-template-columns: repeat(3, minmax(350px, 1fr));
    }
}
</style>
