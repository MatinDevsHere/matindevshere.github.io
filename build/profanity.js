// Function to toggle profanity state
function toggleProfanity() {
    const nsfwElements = document.querySelectorAll('.nsfw');
    const sfwElements = document.querySelectorAll('.sfw');
    const toggleButton = document.getElementById('profanityToggle');
    const currentState = localStorage.getItem('profanityEnabled') === 'true';

    // Toggle state
    const newState = !currentState;
    localStorage.setItem('profanityEnabled', newState);

    // Update button text
    toggleButton.textContent = newState ? 'Enable Safe Mode' : 'Enable Adult Content';

    // Toggle visibility
    nsfwElements.forEach(el => {
        el.style.display = newState ? 'inline' : 'none';
    });

    sfwElements.forEach(el => {
        el.style.display = newState ? 'none' : 'inline';
    });
}

// Initialize profanity state on page load
document.addEventListener('DOMContentLoaded', () => {
    const profanityEnabled = localStorage.getItem('profanityEnabled') === 'true';
    const nsfwElements = document.querySelectorAll('.nsfw');
    const sfwElements = document.querySelectorAll('.sfw');
    const toggleButton = document.getElementById('profanityToggle');

    // Set initial button text
    toggleButton.textContent = profanityEnabled ? 'Enable Safe Mode' : 'Enable Adult Content';

    // Set initial visibility
    nsfwElements.forEach(el => {
        el.style.display = profanityEnabled ? 'inline' : 'none';
    });

    sfwElements.forEach(el => {
        el.style.display = profanityEnabled ? 'none' : 'inline';
    });
}); 