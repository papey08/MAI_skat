using UnityEngine;

public class PlaySoundOnDestroy : MonoBehaviour
{
    public AudioClip destroyClip;
    private static AudioSource globalAudioSource;

    void Start()
    {
        if (globalAudioSource == null)
        {
            GameObject audioObj = new GameObject("GlobalAudioSource");
            globalAudioSource = audioObj.AddComponent<AudioSource>();
            globalAudioSource.loop = false;
            globalAudioSource.playOnAwake = false;
            DontDestroyOnLoad(audioObj);
        }    
    }

    void OnDestroy()
    {
        if (destroyClip != null)
        {
            globalAudioSource.clip = destroyClip;
            globalAudioSource.Play();
        }
    }
}
