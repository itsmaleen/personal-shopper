<template>
  <div class="h-full flex">
    <!-- Narrow sidebar -->
    <div class="hidden w-28 bg-indigo-700 overflow-y-auto md:block">
      <div class="w-full py-6 flex flex-col items-center">
        <div class="flex-shrink-0 flex items-center">
          <img
            class="h-8 w-auto"
            src="https://tailwindui.com/img/logos/workflow-mark.svg?color=white"
            alt="Workflow"
          />
        </div>
        <div class="flex-1 mt-6 w-full px-2 space-y-1">
          <a
            v-for="item in sidebarNavigation"
            :key="item.name"
            :href="item.href"
            :class="[
              item.current
                ? 'bg-indigo-800 text-white'
                : 'text-indigo-100 hover:bg-indigo-800 hover:text-white',
              'group w-full p-3 rounded-md flex flex-col items-center text-xs font-medium',
            ]"
            :aria-current="item.current ? 'page' : undefined"
          >
            <component
              :is="item.icon"
              :class="[
                item.current
                  ? 'text-white'
                  : 'text-indigo-300 group-hover:text-white',
                'h-6 w-6',
              ]"
              aria-hidden="true"
            />
            <span class="mt-2">{{ item.name }}</span>
          </a>
        </div>
      </div>
    </div>

    <!-- Mobile menu -->
    <TransitionRoot as="template" :show="mobileMenuOpen">
      <Dialog as="div" class="md:hidden" @close="mobileMenuOpen = false">
        <div class="fixed inset-0 z-40 flex">
          <TransitionChild
            as="template"
            enter="transition-opacity ease-linear duration-300"
            enter-from="opacity-0"
            enter-to="opacity-100"
            leave="transition-opacity ease-linear duration-300"
            leave-from="opacity-100"
            leave-to="opacity-0"
          >
            <DialogOverlay class="fixed inset-0 bg-gray-600 bg-opacity-75" />
          </TransitionChild>
          <TransitionChild
            as="template"
            enter="transition ease-in-out duration-300 transform"
            enter-from="-translate-x-full"
            enter-to="translate-x-0"
            leave="transition ease-in-out duration-300 transform"
            leave-from="translate-x-0"
            leave-to="-translate-x-full"
          >
            <div
              class="relative max-w-xs w-full bg-indigo-700 pt-5 pb-4 flex-1 flex flex-col"
            >
              <TransitionChild
                as="template"
                enter="ease-in-out duration-300"
                enter-from="opacity-0"
                enter-to="opacity-100"
                leave="ease-in-out duration-300"
                leave-from="opacity-100"
                leave-to="opacity-0"
              >
                <div class="absolute top-1 right-0 -mr-14 p-1">
                  <button
                    type="button"
                    class="h-12 w-12 rounded-full flex items-center justify-center focus:outline-none focus:ring-2 focus:ring-white"
                    @click="mobileMenuOpen = false"
                  >
                    <XIcon class="h-6 w-6 text-white" aria-hidden="true" />
                    <span class="sr-only">Close sidebar</span>
                  </button>
                </div>
              </TransitionChild>
              <div class="flex-shrink-0 px-4 flex items-center">
                <img
                  class="h-8 w-auto"
                  src="https://tailwindui.com/img/logos/workflow-mark.svg?color=white"
                  alt="Workflow"
                />
              </div>
              <div class="mt-5 flex-1 h-0 px-2 overflow-y-auto">
                <nav class="h-full flex flex-col">
                  <div class="space-y-1">
                    <a
                      v-for="item in sidebarNavigation"
                      :key="item.name"
                      :href="item.href"
                      :class="[
                        item.current
                          ? 'bg-indigo-800 text-white'
                          : 'text-indigo-100 hover:bg-indigo-800 hover:text-white',
                        'group py-2 px-3 rounded-md flex items-center text-sm font-medium',
                      ]"
                      :aria-current="item.current ? 'page' : undefined"
                    >
                      <component
                        :is="item.icon"
                        :class="[
                          item.current
                            ? 'text-white'
                            : 'text-indigo-300 group-hover:text-white',
                          'mr-3 h-6 w-6',
                        ]"
                        aria-hidden="true"
                      />
                      <span>{{ item.name }}</span>
                    </a>
                  </div>
                </nav>
              </div>
            </div>
          </TransitionChild>
          <div class="flex-shrink-0 w-14" aria-hidden="true">
            <!-- Dummy element to force sidebar to shrink to fit close icon -->
          </div>
        </div>
      </Dialog>
    </TransitionRoot>

    <!-- Content area -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="w-full">
        <div
          class="relative z-10 flex-shrink-0 h-16 bg-white border-b border-gray-200 shadow-sm flex"
        >
          <button
            type="button"
            class="border-r border-gray-200 px-4 text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500 md:hidden"
            @click="mobileMenuOpen = true"
          >
            <span class="sr-only">Open sidebar</span>
            <MenuAlt2Icon class="h-6 w-6" aria-hidden="true" />
          </button>
          <div class="flex-1 flex justify-between px-4 sm:px-6">
            <div class="flex-1 flex">
              <form class="w-full flex md:ml-0" action="#" method="GET">
                <label for="search-field" class="sr-only"
                  >Search all files</label
                >
                <div
                  class="relative w-full text-gray-400 focus-within:text-gray-600"
                >
                  <div
                    class="pointer-events-none absolute inset-y-0 left-0 flex items-center"
                  >
                    <SearchIcon
                      class="flex-shrink-0 h-5 w-5"
                      aria-hidden="true"
                    />
                  </div>
                  <input
                    name="search-field"
                    id="search-field"
                    class="h-full w-full border-transparent py-2 pl-8 pr-3 text-base text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-0 focus:border-transparent focus:placeholder-gray-400"
                    placeholder="Search"
                    type="search"
                  />
                </div>
              </form>
            </div>
            <div class="ml-2 flex items-center space-x-4 sm:ml-6 sm:space-x-6">
              <!-- Profile dropdown -->
              <Menu as="div" class="relative flex-shrink-0">
                <div>
                  <MenuButton
                    class="bg-white rounded-full flex text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  >
                    <span class="sr-only">Open user menu</span>
                    <img
                      class="h-8 w-8 rounded-full"
                      src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=256&h=256&q=80"
                      alt=""
                    />
                  </MenuButton>
                </div>
                <transition
                  enter-active-class="transition ease-out duration-100"
                  enter-from-class="transform opacity-0 scale-95"
                  enter-to-class="transform opacity-100 scale-100"
                  leave-active-class="transition ease-in duration-75"
                  leave-from-class="transform opacity-100 scale-100"
                  leave-to-class="transform opacity-0 scale-95"
                >
                  <MenuItems
                    class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                  >
                    <MenuItem
                      v-for="item in userNavigation"
                      :key="item.name"
                      v-slot="{ active }"
                    >
                      <a
                        :href="item.href"
                        :class="[
                          active ? 'bg-gray-100' : '',
                          'block px-4 py-2 text-sm text-gray-700',
                        ]"
                        >{{ item.name }}</a
                      >
                    </MenuItem>
                  </MenuItems>
                </transition>
              </Menu>

              <button
                type="button"
                class="flex bg-indigo-600 p-1 rounded-full items-center justify-center text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                <PlusSmIcon class="h-6 w-6" aria-hidden="true" />
                <span class="sr-only">Add file</span>
              </button>
            </div>
          </div>
        </div>
      </header>

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
                    :id="`product-${product.id}`"
                    class="group relative"
                    ref="productsRef"
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
                          <div>No tags</div>
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
          class="hidden w-96 bg-white border-l border-gray-200 overflow-y-auto lg:block"
        >
          <span
            v-for="(tag, index) in tags"
            :key="index"
            class="px-2 py-1 text-green-800 text-xs font-medium bg-green-100 rounded-full cursor-move tag"
            :id="`tag-${index}`"
            draggable="true"
          >
            {{ tag.name }}
            <button
              type="button"
              class="rounded-md p-1.5 text-green-500 focus:outline-none"
              @click="deleteTag"
            >
              <span class="sr-only">Dismiss</span>
              <XIcon class="h-3 w-3" aria-hidden="true" />
            </button>
          </span>
        </aside>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import {
  Dialog,
  DialogOverlay,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  TransitionChild,
  TransitionRoot,
} from "@headlessui/vue";
import {
  CogIcon,
  CollectionIcon,
  HomeIcon,
  MenuAlt2Icon,
  PhotographIcon,
  PlusSmIcon,
  UserGroupIcon,
  ViewGridIcon,
  XIcon,
} from "@heroicons/vue/outline";
import { SearchIcon } from "@heroicons/vue/solid";

const tags = [
  {
    name: "streetwear",
    color: "#312349",
  },
  {
    name: "high fashion",
    color: "#312349",
  },
];

const products = [
  {
    id: 1,
    name: "Basic Tee",
    href: "#",
    imageSrc:
      "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
    imageAlt: "Front of men's Basic Tee in black.",
    price: "$35",
    color: "Black",
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
  },
];

const sidebarNavigation = [
  { name: "Home", href: "#", icon: HomeIcon, current: false },
  { name: "All Files", href: "#", icon: ViewGridIcon, current: false },
  { name: "Photos", href: "#", icon: PhotographIcon, current: true },
  { name: "Shared", href: "#", icon: UserGroupIcon, current: false },
  { name: "Albums", href: "#", icon: CollectionIcon, current: false },
  { name: "Settings", href: "#", icon: CogIcon, current: false },
];
const userNavigation = [
  { name: "Your Profile", href: "#" },
  { name: "Sign out", href: "#" },
];

function getTagTarget(e) {
  return e.target.closest(".tag");
}

function getProductTarget(e) {
  return e.target.closest(".group");
}

function dragStart(e) {
  e.dataTransfer.setData("text/plain", e.target.id);
}

function dragEnter(e) {
  e.preventDefault();
  getProductTarget(e).classList.add("drag-over");
}

function dragOver(e) {
  e.preventDefault();
  getProductTarget(e).classList.add("drag-over");
}

function dragLeave(e) {
  getProductTarget(e).classList.remove("drag-over");
}

function drop(e) {
  const productTarget = getProductTarget(e);
  productTarget.classList.remove("drag-over");

  // get the draggable element
  const id = e.dataTransfer.getData("text/plain");
  const draggable = document.getElementById(id);

  // add it to the drop target
  let tagAlreadyExist = false;
  for (let i = 0; i < productTarget.children.length; i++) {
    if (productTarget.children[i].textContent === draggable.textContent) {
      tagAlreadyExist = true;
      // TODO: Add toast that says tag already added
      break;
    }
  }
  if (!tagAlreadyExist) {
    console.log();
    const clone = draggable.cloneNode(true);
    clone.draggable = false;
    clone.classList.remove("cursor-move");
    clone.classList.add("cursor-default");
    clone.childNodes[1].classList.remove("hidden");
    clone.id = `${productTarget.id}-${draggable.id}`;
    productTarget.appendChild(clone);
  }
}

export default {
  components: {
    Dialog,
    DialogOverlay,
    Menu,
    MenuButton,
    MenuItem,
    MenuItems,
    TransitionChild,
    TransitionRoot,
    MenuAlt2Icon,
    PlusSmIcon,
    SearchIcon,
    XIcon,
  },
  setup() {
    const mobileMenuOpen = ref(false);

    return {
      tags,
      products,
      sidebarNavigation,
      userNavigation,
      mobileMenuOpen,
    };
  },
  mounted() {
    tags.forEach((tag, index) => {
      const tagElement = document.getElementById(`tag-${index}`);
      tagElement.addEventListener("dragstart", dragStart);
    });
    products.forEach((product) => {
      const productElement = document.getElementById(`product-${product.id}`);
      productElement.addEventListener("dragenter", dragEnter);
      productElement.addEventListener("dragover", dragOver);
      productElement.addEventListener("dragleave", dragLeave);
      productElement.addEventListener("drop", drop);
    });
  },
  methods: {
    hasTags(index) {
      console.log(this.$refs);
      if (this.$refs?.productsRef) {
        console.log(this.$refs.productsRef);
        return this.$refs?.productsRef[index]?.children.length > 2;
      }
      return false;
    },
    deleteTag(event) {
      console.log("delete tag");
      const tagElement = getTagTarget(event);
      tagElement.removeEventListener("dragstart", dragStart);
      tagElement.removeEventListener("dragenter", dragEnter);
      tagElement.removeEventListener("dragover", dragOver);
      tagElement.removeEventListener("dragleave", dragLeave);
      tagElement.removeEventListener("drop", drop);
      tagElement.remove();
    },
  },
};
</script>

<style>
.drag-over {
  @apply border-dashed border-4 border-gray-300;
}
</style>
