using UnityEngine;

[RequireComponent(typeof(Renderer))]
public class MusicTriggerOnVisible : MonoBehaviour
{
    public AudioClip musicClip;
    private AudioSource musicSource;

    private void Start()
    {
        musicSource = gameObject.AddComponent<AudioSource>();
        musicSource.clip = musicClip;
        musicSource.loop = true;
        musicSource.playOnAwake = false;
        musicSource.volume = 0.7f;
    }

    private void OnBecameVisible()
    {
        if (!musicSource.isPlaying)
        {
            musicSource.Play();
        }
    }

    private void OnBecameInvisible()
    {
        if (musicSource.isPlaying)
        {
            musicSource.Stop();
        }
    }
}
