<template>
  <!-- Main content -->
  <div class="flex-1 flex items-stretch overflow-hidden">
    <main class="flex-1 overflow-y-auto">
      <!-- Primary column -->
      <section
        aria-labelledby="primary-heading"
        class="min-w-0 flex-1 h-full flex flex-col lg:order-last"
      >
        <h1 id="primary-heading" class="sr-only">Photos</h1>
        <div class="bg-white">
          <div
            class="max-w-2xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8"
          >
            <div
              class="mt-6 grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8"
            >
              <div
                v-for="product in products"
                :key="product.id"
                class="group relative"
                v-on:dragenter="dragEnter($event, product)"
                v-on:dragover="dragOver($event, product)"
                v-on:dragleave="dragLeave($event, product)"
                v-on:drop="drop($event, product)"
              >
                <div
                  class="w-full min-h-80 bg-gray-200 aspect-w-1 aspect-h-1 rounded-md overflow-hidden group-hover:opacity-75 lg:h-80 lg:aspect-none"
                >
                  <img
                    :src="product.imageSrc"
                    :alt="product.imageAlt"
                    class="w-full h-full object-center object-cover lg:w-full lg:h-full"
                  />
                </div>
                <div class="mt-4 flex justify-between">
                  <div>
                    <h3 class="text-sm text-gray-700">
                      <a :href="product.href">
                        {{ product.name }}
                      </a>
                    </h3>
                    <!-- Hide tags -->
                    <div class="mt-1 text-sm text-gray-500">
                      <div v-if="product.tags.length === 0">No tags</div>
                      <div v-else>
                        <span
                          v-for="(tag, index) in product.tags"
                          :key="index"
                          class="px-2 py-1 text-green-800 text-xs font-medium bg-green-100 rounded-full cursor-move tag"
                        >
                          {{ tag.name }}
                          <button
                            type="button"
                            class="rounded-md p-1.5 text-green-500 focus:outline-none"
                            @click="removeTag(product.id, tag)"
                          >
                            <span class="sr-only">Dismiss</span>
                            <XIcon class="h-3 w-3" aria-hidden="true" />
                          </button>
                        </span>
                      </div>
                    </div>
                  </div>
                  <p class="text-sm font-medium text-gray-900">
                    {{ product.price }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>

    <!-- Secondary column (hidden on smaller screens) -->
    <aside
      class="hidden w-96 bg-white border-l border-gray-200 overflow-y-auto lg:block p-8"
    >
      <!-- <span
            v-for="(tag, index) in tags"
            :key="index"
            class="px-2 py-1 text-green-800 text-xs font-medium bg-green-100 rounded-full cursor-move tag"
            :id="`tag-${index}`"
            draggable="true"
          >
            {{ tag.name }}
          </span> -->
      <ul
        role="list"
        class="mt-2 border-t border-b border-gray-200 divide-y divide-gray-200"
      >
        <li
          v-for="(tag, index) in tags"
          :key="index"
          class="py-3 flex justify-between items-center"
          draggable="true"
          v-on:dragstart="dragStart($event, tag)"
        >
          <div class="flex items-center">
            <div class="w-8 h-8 rounded-full bg-green-100" />
            <p class="ml-4 text-sm font-medium text-gray-900">
              {{ tag.name }}
            </p>
          </div>
          <button
            type="button"
            class="ml-6 bg-white rounded-md text-sm font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Remove<span class="sr-only"> {{ tag.name }}</span>
          </button>
        </li>
        <li class="py-2 flex justify-between items-center">
          <label for="tag" class="sr-only">tag</label>
          <input
            type="text"
            name="tag"
            id="tag"
            class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
            placeholder="streetwear"
            v-model="newTag"
          />
          <button
            type="button"
            @click="createTag"
            class="ml-6 bg-white rounded-md text-sm font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Add<span class="sr-only"> Label</span>
          </button>
        </li>
      </ul>
    </aside>
  </div>
</template>

<script>
import { ref } from "vue";
import { XIcon } from "@heroicons/vue/outline";
import gql from "graphql-tag";

function getProductTarget(e) {
  return e.target.closest(".group");
}

export default {
  components: {
    XIcon,
  },
  apollo: {
    tags: gql`
      query {
        tags {
          id
          name
        }
      }
    `,
  },
  data() {
    return {
      products: [
        {
          id: 1,
          name: "Basic Tee",
          href: "#",
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
          imageAlt: "Front of men's Basic Tee in black.",
          price: "$35",
          color: "Black",
          tags: [],
        },
        {
          id: 2,
          name: "Basic Tee",
          href: "#",
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
          imageAlt: "Front of men's Basic Tee in black.",
          price: "$35",
          color: "Black",
          tags: [],
        },
        {
          id: 3,
          name: "Basic Tee",
          href: "#",
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
          imageAlt: "Front of men's Basic Tee in black.",
          price: "$35",
          color: "Black",
          tags: [],
        },
      ],
      newTag: null,
    };
  },
  methods: {
    dragStart(e, tag) {
      console.log("dragstart ", tag);
      e.dataTransfer.setData("text/json", JSON.stringify(tag));
    },
    dragEnter(e) {
      e.preventDefault();
      getProductTarget(e).classList.add("drag-over");
    },
    dragOver(e) {
      e.preventDefault();
      getProductTarget(e).classList.add("drag-over");
    },
    dragLeave(e) {
      getProductTarget(e).classList.remove("drag-over");
    },
    drop(e, product) {
      getProductTarget(e).classList.remove("drag-over");

      const tagJson = e.dataTransfer.getData("text/json");
      const tag = JSON.parse(tagJson);

      const tagIndex = product.tags.findIndex((t) => t.name === tag.name);

      if (tagIndex === -1) {
        product.tags.push({ name: tag.name });
      }
    },
    removeTag(productId, tag) {
      const product = this.products.find((product) => product.id === productId);
      const tagIndex = product.tags.findIndex((t) => t.name === tag.name);
      product.tags.splice(tagIndex, 1);
    },
    async createTag() {
      if (this.newTag) {
        const result = await this.$apollo.mutate({
          mutation: gql`
            mutation createTag($newTag: NewTag!) {
              createTag(input: $newTag) {
                id
                name
              }
            }
          `,
          variables: {
            newTag: {
              name: this.newTag,
              color: "#D960AA",
            },
          },
        });
        this.newTag = null;
      }
    },
  },
};
</script>

<style>
.drag-over {
  @apply border-dashed border-4 border-gray-300;
}
</style>
