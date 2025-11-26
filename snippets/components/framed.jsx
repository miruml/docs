export const Framed = ({
    image,
    background,
    link,
    alt = "Framed content",
    borderWidth = "10px",
    outerRadius = "10px",
    innerRadius = "0",
}) => {
    const innerImage = (
        <img
            src={image}
            alt={alt}
            noZoom={link ? true : false}
            style={{
                display: "block",
                borderRadius: innerRadius,
                margin: `0`,
            }}
        />
    );

    return (
        <div
            style={{
                backgroundImage: background ? `url(${background})` : undefined,
                backgroundSize: "cover",
                borderRadius: outerRadius,
                padding: `${borderWidth}`,
                overflow: "hidden",
            }}
        >
            <div>
                {link ? (<a href={link}> {innerImage} </a>) : (innerImage)}
            </div>
        </div>
    );
};
